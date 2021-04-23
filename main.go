package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Page struct {
	ID             string    `gorm:"primaryKey"`
	Email          string    `gorm:"uniqueIndex;not null"`
	Page           string    `gorm:"type:text;not null"`
	Hidden         string    `gorm:"type:text;not null"`
	LastModifiedBy string    `gorm:"type:text;not null"`
	CreatedAt      time.Time `gorm:"not null"`
	UpdatedAt      time.Time `gorm:"not null"`
}

func (page *Page) Lookup(s string) (string, bool) {
	for _, chunk := range strings.Split(page.Hidden, "|") {
		splitted := strings.SplitN(chunk, "=", 2)
		if len(splitted) == 2 {
			key := splitted[0]
			value := splitted[1]
			if key == s {
				return value, true
			}
		}
	}
	return "", false
}

func valid(s string) bool {
	return len(s) == 8
}

func sanitize(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r
		}
		return -1
	}, strings.ToLower(s))
}

type Config struct {
	CookieSecret string            `json:"cookie_secret"`
	SlackHooks   map[string]string `json:"slack"`
	ClientID     string            `json:"oauth_client_id"`
	ClientSecret string            `json:"oauth_client_secret"`
	PublicCircle string            `json:"public_circle"`
	SendgridKey  string            `json:"sendgrid_key"`
	DB           string            `json:"db"`
}

func ReadConfig(fn string) *Config {
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	c := &Config{}
	err = json.Unmarshal(b, c)
	if err != nil {
		panic(err)
	}

	return c
}
func main() {
	templates := flag.String("templates", "/var/www/detective", "templates")
	local := flag.Bool("local-dev", false, "run locally")
	migrate := flag.Bool("migrate", false, "migrate")
	configPath := flag.String("config", "/etc/detective/config.json", "config path")
	bind := flag.String("bind", "127.0.0.1:8080", "bind")

	flag.Parse()

	var config *Config
	if *local {
		config = &Config{
			PublicCircle: "dance",
			SlackHooks:   map[string]string{},
			CookieSecret: "hello",
		}
	} else {
		config = ReadConfig(*configPath)
	}

	var sqlDB gorm.Dialector
	if *local {
		sqlDB = sqlite.Open("test.db")
	} else {
		sqlDB = postgres.Open(config.DB)
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	db, err := gorm.Open(sqlDB, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	if *migrate {
		db.AutoMigrate(&Page{})
	}

	r := gin.Default()

	store := sessions.NewCookieStore([]byte(config.CookieSecret))

	scopes := []string{
		"https://www.googleapis.com/auth/userinfo.email",
	}

	r.Use(sessions.Sessions("detective", store))
	if *local {
		r.Use(FakeAuth(config.ClientSecret, config.ClientID, scopes))
	} else {
		r.Use(Auth(config.ClientSecret, config.ClientID, scopes))
	}

	// NB: this is also public accessible
	r.LoadHTMLGlob(path.Join(*templates, "*.html"))
	r.Static("/static", path.Join(*templates, "..", "static"))
	r.GET("/auth", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	r.GET("/x/:id/:key", func(c *gin.Context) {
		c.Header("X-Robots-Tag", "noindex")
		id := sanitize(c.Param("id"))
		if !valid(id) {
			c.String(400, "detective agency name has to be exactly 8 characters")
			return
		}

		page := &Page{}
		if err := db.Where("id = ?", id).First(page).Error; err != nil {
			panic(err)
		}
		value, ok := page.Lookup(c.Param("key"))
		if ok {
			c.String(200, value)
			return
		}
		c.String(404, "key_not_found")
	})

	r.GET("/d/:id", func(c *gin.Context) {
		c.Header("X-Robots-Tag", "noindex")

		id := sanitize(c.Param("id"))
		if !valid(id) {
			c.String(400, "detective agency name has to be exactly 8 characters")
			return
		}

		page := &Page{}
		if err := db.Where("id = ?", id).First(page).Error; err != nil {
			c.String(200, fmt.Sprintf("/d/%s is unclaimed, go to /register to make it your page", id))
			return
		}
		c.Data(200, "text/html; charset=utf-8", []byte(page.Page))
	})

	r.GET("/logout", func(c *gin.Context) {
		s := sessions.Default(c)
		s.Clear()
		s.Save()

		c.Redirect(302, "/")
	})

	r.GET("/register", func(c *gin.Context) {
		email := c.MustGet("email").(string)
		page := &Page{}

		if err := db.Where("email = ?", email).First(page).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				panic(err)
			}
		} else {
			c.Redirect(302, "/edit")
			return
		}
		c.HTML(200, "register.html", gin.H{"email": email})
	})

	r.POST("/register", func(c *gin.Context) {
		id := sanitize(c.PostForm("page"))
		if !valid(id) {
			c.String(400, "detective agency name has to be exactly 8 characters")
			return
		}

		page := &Page{}
		tx := db.Begin()
		defer func() {
			tx.Rollback()
			if r := recover(); r != nil {
				panic(r)
			}
		}()

		if err := tx.Where("id = ?", id).First(page).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				panic(err)
			}
		} else {
			c.String(200, "someone claimed this page before you")
			return
		}

		page.ID = id
		page.Email = c.MustGet("email").(string)
		page.Hidden = ""
		page.Page = fmt.Sprintf(`<html>
<body>
	<h2>welcome!</h2>
	<p>
		this is my <b>first web page</b>!
	</p>
	and it has a bug!<br>
	<button onclick="document.body.appendChild(this.cloneNode(true))">🐛</button>
</body>
</html>`)

		dump, err := httputil.DumpRequest(c.Request, false)
		if err != nil {
			log.Printf("error dumping http request, err: %v", err)
		} else {
			if len(dump) > 65000 {
				dump = dump[:65000]
			}
			page.LastModifiedBy = string(dump)
		}

		if err := tx.Create(page).Error; err != nil {
			panic(err)
		}

		if err := tx.Commit().Error; err != nil {
			panic(err)
		}

		c.Redirect(302, "/edit")
	})

	admin := func(c *gin.Context) {
		email := c.MustGet("email").(string)

		page := &Page{}
		if err := db.Where("email = ?", email).First(page).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.Redirect(302, "/register")
				return
			}
			panic(err)
		}

		if c.Request.Method == "POST" {
			page.Page = c.PostForm("page")
			page.Hidden = c.PostForm("hidden")

			if len(page.Page) > 4096 || len(page.Hidden) > 4096 {
				c.String(200, "a detective agency page or hidden text can be at most 4096 characters! go back and try again.")
				return
			}
			dump, err := httputil.DumpRequest(c.Request, false)
			if err != nil {
				log.Printf("error dumping http request, err: %v", err)
			} else {
				if len(dump) > 65000 {
					dump = dump[:65000]
				}
				page.LastModifiedBy = string(dump)

			}

			if err := db.Save(page).Error; err != nil {
				panic(err)
			}
		}

		c.HTML(http.StatusOK, "edit.html", gin.H{"id": page.ID, "page": page, "email": email, "showPreview": c.Query("show_preview") != "false"})
	}
	r.GET("/edit", admin)
	r.POST("/edit", admin)
	r.Run(*bind)
}
