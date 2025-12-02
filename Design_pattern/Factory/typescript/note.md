# Factory Pattern ใน Go และ TypeScript

## ความสำคัญใน Go (Golang)
Go จะเห็น pattern ชัดมาก ๆ เพราะ **Go ไม่มี constructor แบบ class** แต่แพทเทิร์นนี้ใช้เยอะสุด ๆ ใน service backend.

### ก่อนใช้ Factory
- ทุกไฟล์ต้องรู้ว่าเราจะ `new` อะไร → ทำให้ระบบ **coupling สูง**
- แก้หลายที่ และต้องมี **if-else กระจาย**

### หลังใช้ Factory
- ฟังก์ชันทุกที่จะใช้ผ่านการ `create` จาก Factory เดียว
- **เพิ่ม product ใหม่ไม่ต้องแก้ client code**
- ควบคุม dependency ได้จากจุดเดียว

## Simple Factory คืออะไร?
**Simple Factory** = รวม logic การสร้าง object ที่ซ้ำ ๆ ไว้ที่เดียว → เพื่อกระจายอำนาจการเลือก/สร้าง object ไปจาก client code

หรือพูดให้สวยในเชิงสถาปัตยกรรม:
- **Centralize object creation**
- ลด duplication
- ลด coupling
- ทำให้ client “ไม่ต้องรู้” ว่าต้อง `new` อะไร
- เปลี่ยน implementation ได้ที่เดียว

ง่ายแต่ทรงพลังมาก!


## Diagram
![Factory Pattern Diagram](Method_vs_Abstract.png)


## Concrete Classes
```typescript
// import { CreditCardPayment } from "./payments/CreditCardPayment";
// import { PaypalPayment } from "./payments/PaypalPayment";
// import { PromptPayPayment } from "./payments/PromptPayPayment";

