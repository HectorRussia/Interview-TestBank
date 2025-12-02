interface ShippingStrategy {
  calculate(cost: number): number;
}

class StandardShipping implements ShippingStrategy {
  calculate(cost: number): number {
    return cost + 50;
  }
}

class ExpressShipping implements ShippingStrategy {
  calculate(cost: number): number {
    return cost + 100;
  }
}

class Order {
  constructor(private strategy: ShippingStrategy) {}

  getTotal(price: number): number {
    return this.strategy.calculate(price);
  }
}

// การใช้งาน
const order1 = new Order(new StandardShipping());
console.log(order1.getTotal(500)); // 550

const order2 = new Order(new ExpressShipping());
console.log(order2.getTotal(500)); // 600

// เปลี่ยน “พฤติกรรม / อัลกอริทึม” โดยไม่ต้อง if-else ยาวเป็นหางว่าว

// ใช้เมื่อไหร่

// มี logic หลายแบบ เช่น

// คำนวณค่าขนส่งแบบต่าง ๆ

// promotion, discount หลายสูตร

// ไม่อยากให้ class หลักเต็มไปด้วย if (type === "...")