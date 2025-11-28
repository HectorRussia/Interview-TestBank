class Logger {
    
    private static instance: Logger | null = null;

    private constructor() {} //  box new outside

    static getInstance(): Logger {
        if(!Logger.instance) {
            Logger.instance = new Logger();
        }
        return Logger.instance
    }

    log(message: string) {
        console.log("this message is" , message);
    }
}

// use case
const loger1 = Logger.getInstance()
const loger2 = Logger.getInstance()

console.log(loger1 === loger2);

// Singleton Pattern
// มีแค่ “หนึ่ง instance” ทั้งระบบ เช่น Logger, Config, DB connection

// ใช้เมื่อไหร่

// ของชิ้นนั้นต้องมีแค่ตัวเดียว เช่น

// Global config

// Connection pool

// Central logger

// ประเด็นที่ชอบถาม

// ข้อดี: control instance, share state

// ข้อเสีย: test ยาก, global state, ถ้าใช้มั่วจะเละ