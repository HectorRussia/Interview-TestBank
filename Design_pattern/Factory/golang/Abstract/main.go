package main

import "fmt"

// ----- Products -----
type Button interface {
	Render()
}

type Checkbox interface {
	Render()
}

// Windows products
type WinButton struct{}

func (WinButton) Render() {
	fmt.Println("Render Windows Button")
}

type WinCheckbox struct{}

func (WinCheckbox) Render() {
	fmt.Println("Render Windows Checkbox")
}

// Mac products
type MacButton struct{}

func (MacButton) Render() {
	fmt.Println("Render Mac Button")
}

type MacCheckbox struct{}

func (MacCheckbox) Render() {
	fmt.Println("Render Mac Checkbox")
}

// ----- Abstract Factory -----
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// ----- Concrete Factories -----
type WinFactory struct{}

func (WinFactory) CreateButton() Button {
	return WinButton{}
}

func (WinFactory) CreateCheckbox() Checkbox {
	return WinCheckbox{}
}

type MacFactory struct{}

func (MacFactory) CreateButton() Button {
	return MacButton{}
}

func (MacFactory) CreateCheckbox() Checkbox {
	return MacCheckbox{}
}

// ----- Client -----
type Application struct {
	button   Button
	checkbox Checkbox
}

func NewApplication(factory GUIFactory) *Application {
	return &Application{
		button:   factory.CreateButton(),
		checkbox: factory.CreateCheckbox(),
	}
}

func (a *Application) RenderUI() {
	a.button.Render()
	a.checkbox.Render()
}

func main() {
	var factory GUIFactory

	// สมมติ detect OS จาก config
	os := "mac"

	if os == "windows" {
		factory = WinFactory{}
	} else {
		factory = MacFactory{}
	}

	app := NewApplication(factory)
	app.RenderUI()

	// ถ้า os = "mac" → Render Mac Button + Mac Checkbox
	// ถ้า os = "windows" → Render Windows Button + Windows Checkbox
}

// Keyword:
// โรงงานที่สร้าง “หลาย product ที่เป็น family เดียวกัน” พร้อมกัน
// เช่น ตระกูล UI สำหรับ Windows / Mac

// สถานการณ์: สร้าง UI ตาม OS

// Product family:

// Button

// Checkbox

// Theme:

// Windows

// Mac

// Key จุดต่างจาก Factory Method

// Factory Method → เน้น “เมธอดเดียว” ใน Creator สร้าง product เดียว

// Abstract Factory → มี interface โรงงาน (GUIFactory) ที่มี “หลาย factory method”
// (CreateButton, CreateCheckbox, ...)

// และสิ่งสำคัญของ Abstract Factory คือ:

// แถม guarantee ว่า objects ที่สร้าง “เข้าชุดกัน” (Mac button + Mac checkbox, ไม่ปน Win button + Mac checkbox)
