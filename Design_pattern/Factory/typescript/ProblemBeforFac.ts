class CreditCardPayment  {
    pay(amount: number): void {
        console.log("Paying", amount , "with Credit Cards");
    }
}

class PromptPayPayment  {
    pay(amount: number): void {
        console.log("Paying", amount , "with Crash");
    }
}

class PaypalPayment  {
   pay(amount: number): void {
        console.log("Paying", amount , "with Paypal");
   }
}

// assume class us live another there
// concrete class
// import { CreditCardPayment } from "./payments/CreditCardPayment";
// import { PaypalPayment } from "./payments/PaypalPayment";
// import { PromptPayPayment } from "./payments/PromptPayPayment";

// ก็ไม่ได้แย่ แต่ถ้ามีฟังก์ชั่นแบบนี้เยอะๆละ ?
function checkout(amount: number, method: string) {
  if (method === "card") {
    const payment = new CreditCardPayment();
    payment.pay(amount);
  } else if (method === "paypal") {
    const payment = new PaypalPayment();
    payment.pay(amount);
  } else if (method === "promptpay") {
    const payment = new PromptPayPayment();
    payment.pay(amount);
  }
}

function renewSubscription(amount: number, method: string) {
  if (method === "card") {
    const payment = new CreditCardPayment();
    payment.pay(amount);
  } else if (method === "paypal") {
    const payment = new PaypalPayment();
    payment.pay(amount);
  } else if (method === "promptpay") {
    const payment = new PromptPayPayment();
    payment.pay(amount);
  }
}


function buyCourse(amount: number, method: string) {
  if (method === "card") {
    const payment = new CreditCardPayment();
    payment.pay(amount);
  } else if (method === "paypal") {
    const payment = new PaypalPayment();
    payment.pay(amount);
  } else if (method === "promptpay") {
    const payment = new PromptPayPayment();
    payment.pay(amount);
  }
}


// อีกปัญหา
// user บอก อยากจ่ายเงินวิธีใหม่จุง crypto

/* else if (method === "crypto") {
  const payment = new CryptoPayment();
  payment.pay(amount);
} */

// แก้ 1 ที่? ไม่ 

// ต้องแก้ ทุกไฟล์ ที่ new Payment อยู่ → อาจกระจาย 5, 10, 20 จุด!

// ทำให้โค้ดทุกส่วน “รู้จัก concrete class”

// ทุกฟังก์ชัน import แบบนี้:

// import { CreditCardPayment } from "./CreditCardPayment";
// import { PaypalPayment } from "./PaypalPayment";
// ทำให้ ผูก dependency แน่นมาก
// เปลี่ยนชื่อไฟล์ เปลี่ยนโฟลเดอร์ เปลี่ยน constructor → พังทั้งระบบ

// if-else กระจายเต็มระบบ

// เปลี่ยน logic ที่เลือก Payment ต้องไล่แก้ทุกที่:

// if (country === "TH") use PromptPay
// if (country === "US") use PayPal
// if (merchantType === "subscription") use CreditCard


