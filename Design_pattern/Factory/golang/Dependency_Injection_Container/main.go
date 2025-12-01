package main

import (
	"fmt"
	"log"
)

// ---- Low Level Dependency ----
type DB struct {
	DSN string
}

// ---- Domain Layer ----
type UserRepository interface {
	FindNameByID(id int) (string, error)
}

type userRepositoryImpl struct {
	db *DB
}

func (r *userRepositoryImpl) FindNameByID(id int) (string, error) {
	// สมมติอ่านจาก DB จริง ๆ
	return fmt.Sprintf("user-%d-from-%s", id, r.db.DSN), nil
}

// Service
type UserService interface {
	GetUserName(id int) (string, error)
}

type userServiceImpl struct {
	repo UserRepository
}

func (s *userServiceImpl) GetUserName(id int) (string, error) {
	return s.repo.FindNameByID(id)
}

// Handler
type UserHandler struct {
	service UserService
}

func (h *UserHandler) HandleGetUser(id int) {
	name, err := h.service.GetUserName(id)
	if err != nil {
		log.Println("error:", err)
		return
	}
	fmt.Println("User name:", name)
}

// ---- Simple DI Container ----

type Container struct {
	providers map[string]func(*Container) any
	instances map[string]any
}

func NewContainer() *Container {
	return &Container{
		providers: map[string]func(*Container) any{},
		instances: map[string]any{},
	}
}

func (c *Container) Register(name string, provider func(*Container) any) {
	c.providers[name] = provider
}

func (c *Container) Resolve(name string) any {
	// ถ้าเคยสร้างแล้ว ใช้ instance เดิม (singleton-style)
	if inst, ok := c.instances[name]; ok {
		return inst
	}

	provider, ok := c.providers[name]
	if !ok {
		panic("no provider for: " + name)
	}

	inst := provider(c)
	c.instances[name] = inst
	return inst
}

// helper ให้ type-safe ขึ้นนิดหนึ่ง
func Resolve[T any](c *Container, name string) T {
	v := c.Resolve(name)
	return v.(T)
}

// ---- Wiring ทั้งหมดใน main ----

func main() {
	container := NewContainer()

	// Register DB
	container.Register("db", func(c *Container) any {
		return &DB{DSN: "mysql://localhost:3306/app"}
	})

	// Register UserRepository
	container.Register("userRepo", func(c *Container) any {
		db := Resolve[*DB](c, "db")
		return &userRepositoryImpl{db: db}
	})

	// Register UserService
	container.Register("userService", func(c *Container) any {
		repo := Resolve[UserRepository](c, "userRepo")
		return &userServiceImpl{repo: repo}
	})

	// Register UserHandler
	container.Register("userHandler", func(c *Container) any {
		svc := Resolve[UserService](c, "userService")
		return &UserHandler{service: svc}
	})

	// ใช้งาน
	handler := Resolve[*UserHandler](container, "userHandler")
	handler.HandleGetUser(42)
}

// ใน Go ปกติสาย clean code จะชอบ manual DI (ส่ง dependency ผ่าน constructor) มากกว่า container แบบ framework (เหมือน Nest/ Spring)
// แต่เราทำ container เบา ๆ เพื่อให้เห็นภาพ concept IoC/DI ได้

// เป้าหมาย

// มี structure แบบนี้:

// DB – low level dependency

// UserRepository – ใช้ DB

// UserService – ใช้ UserRepository

// UserHandler – ใช้ UserService

// แล้วเราจะมี Container ที่:

// Register constructor ของ service ต่าง ๆ

// Resolve ได้ว่าอยากได้ instance ไหน

// อันนี้จะออกแนว “Service Locator / Simple DI Container” มากกว่าของจริงที่ซับซ้อน แต่พอให้เห็นภาพ

// ตรงนี้ถือเป็น Mini IoC Container

// Container.Register(name, provider) → บอกวิธีสร้าง object นั้น

// Container.Resolve(name) → ขอ instance ของ object นั้น

// Container เก็บ instances ไว้ → ทำตัวเหมือน singleton per service

// ใช้ generic Resolve[T] เพื่อให้ cast ง่ายขึ้น

// ในเชิง concept:

// เรา “กลับด้าน control” การสร้าง object ให้ container เป็นคน control (Inversion of Control)

// โค้ดส่วนอื่นแค่ขอ resolver / handler จาก container

// แต่ในโลก Go จริง ๆ:

// คนจะชอบทำแบบ manual DI มากกว่า (เขียน constructor function และ wire เองใน main.go)

// พวก container แบบนี้ใช้ในโปรเจกต์ใหญ่ ๆ ที่ dependency ซับซ้อนมาก ๆ
