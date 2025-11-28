interface Payment {
    pay(amount: number) : void;
} 

class CreditCardPayment implements Payment {
    pay(amount: number): void {
        console.log("Paying", amount , "with Credit Cards");
    }
}

class CrashPayment implements Payment {
    pay(amount: number): void {
        console.log("Paying", amount , "with Crash");
    }
}

class PayPalPayment implements Payment {
   pay(amount: number): void {
        console.log("Paying", amount , "with Paypal");
   }
}

class PaymentFactory {
    static create(type: "card" | "crash" | "paypal") : Payment {
        if (type === "card") return new CreditCardPayment();
        else if (type === "crash") return new CrashPayment();
        return new PayPalPayment();
    }
}

// use case
const payment = PaymentFactory.create("paypal");
payment.pay(200)


/*  แยก “การสร้าง object” ออกมาอยู่ที่ Factory แทนที่จะ new ตรง ๆ
    ใช้เมื่อไหร่
    มีหลาย class ที่ implement interface เดียวกัน
    อยากให้ client ไม่ต้องรู้รายละเอียดว่า new อะไร 
*/


/*  คำถามที่ชอบเจอ
    Factory ต่างจาก new ปกติยังไง
    Factory Method vs Abstract Factory (ระดับยากขึ้น) 
*/