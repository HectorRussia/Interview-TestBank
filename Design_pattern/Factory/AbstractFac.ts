interface Button {
    render(): void;
}

interface CheckBox {
    render(): void;
}

class WindowButton implements Button {
    render(): void {
        console.log("Render Window Button");
    }
}

class WindowCheckBox implements CheckBox {
    render(): void {
        console.log("Render Mac Checkbox")
    }
}

class MacbookButton implements Button {
    render():void {
        console.log("Render Macbook Button");
    }
}

class MacCheckBox implements CheckBox {
    render(): void {
        console.log("Render Mac Checkbox")
    }
}

// Abstract Factory
interface GUIFactory {
    createButton(): Button;
    createCheckBox(): CheckBox;
}

// Concreate Factories
class WindowFactory implements GUIFactory {
    createButton(): Button {
        return new WindowButton();
    }
    createCheckBox(): CheckBox {
        return new WindowCheckBox();
    }
}

class MacbookFactory implements GUIFactory {

    createButton(): Button {
        return new MacbookButton();
    }

    createCheckBox(): CheckBox {
        return new MacCheckBox();
    }
}

class Application {
    private button: Button;
    private checkBox: CheckBox;

    constructor(factory : GUIFactory) {
        this.button = factory.createCheckBox();
        this.checkBox = factory.createCheckBox();
    }
    
    render() {
        this.button.render()
        this.checkBox.render()
    }
}

const isMac = true;

const factory : GUIFactory = isMac ? new MacbookFactory() : new WindowFactory()

const app = new Application(factory);

app.render();


// Key point ของ Abstract Factory

// มี interface โรงงานเดียว (GUIFactory)
// แต่สร้างได้หลาย product: createButton(), createCheckbox()

// Concrete factory แต่ละอัน (WinFactory, MacFactory) จะสร้าง product ให้เป็นชุดเดียวกัน 
// (Windows theme, Mac theme)

// Client รู้จักแค่ GUIFactory และ interfaces ของ product (Button, Checkbox)