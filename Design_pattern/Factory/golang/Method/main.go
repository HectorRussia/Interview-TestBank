package main

import "fmt"

// Product
type Notification interface {
	Send(message string)
}

// Concrete Products
type EmailNotification struct{}

func (e EmailNotification) Send(message string) {
	fmt.Println("Send EMAIL:", message)
}

type SMSNotification struct{}

func (s SMSNotification) Send(message string) {
	fmt.Println("Send SMS:", message)
}

// Creator
type Notifier interface {
	// Factory Method
	CreateNotification() Notification

	// Template method ที่ใช้ factory method
	SendAlert(msg string)
}

// Concrete Creator 1
type EmailNotifier struct{}

func (e EmailNotifier) CreateNotification() Notification {
	return EmailNotification{}
}

func (e EmailNotifier) SendAlert(msg string) {
	notification := e.CreateNotification() // <= factory method
	notification.Send("[ALERT] " + msg)
}

// Concrete Creator 2
type SMSNotifier struct{}

func (s SMSNotifier) CreateNotification() Notification {
	return SMSNotification{}
}

func (s SMSNotifier) SendAlert(msg string) {
	notification := s.CreateNotification()
	notification.Send("[ALERT] " + msg)
}

func main() {
	var notifier Notifier

	notifier = EmailNotifier{}
	notifier.SendAlert("Server CPU > 90%")
	// Send EMAIL: [ALERT] Server CPU > 90%

	notifier = SMSNotifier{}
	notifier.SendAlert("DB connection failed")
	// Send SMS: [ALERT] DB connection failed
}

// Keyword:
// มี “Creator” ที่มีเมธอดสำหรับสร้าง Product (CreateX) แล้วให้ concrete creator เป็นคนตัดสินใจว่าจะสร้างตัวไหน

// ใน Go ไม่มี inheritance แต่เราใช้ interface + struct แทน “base class / subclass” ได้

// สถานการณ์: Alert Service 2 แบบ

// ส่งแจ้งเตือนผ่าน Email

// ส่งแจ้งเตือนผ่าน SMS

// ใช้ pattern: Application เรียก SendAlert() แต่แต่ละ service จะใช้ Factory Method CreateNotification() ของตัวเอง

// จุดสำคัญของ Factory Method ในโค้ดนี้

// Notifier คือ “Creator interface”

// มี CreateNotification() = Factory Method

// SendAlert() เรียก CreateNotification() แทนที่จะ new เอง

// EmailNotifier / SMSNotifier คือ concrete creator

// ตัดสินใจเองว่าควรสร้าง product ตัวไหน (EmailNotification หรือ SMSNotification)

// Client ใช้ผ่าน Notifier (interface)
// เวลาเปลี่ยนจาก Email → SMS → Line ฯลฯ ก็เปลี่ยน concrete creator ตัวเดียว

// ต่างจาก Simple Factory ตรงที่:

// Simple Factory = ฟังก์ชัน static แบบ NewPayment(method)

// Factory Method = ให้ “ตัว object” (creator) ตัดสินใจเอง ว่าจะสร้าง product แบบไหน โดย override/implement factory method
