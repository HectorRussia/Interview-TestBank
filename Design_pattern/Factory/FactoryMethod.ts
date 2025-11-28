interface Editor {
    open(): void;
}

class VsCode implements Editor {
    open(): void {
        console.log("Open Vscode!");
    }
}

class Cursor implements Editor {
    open(): void {
        console.log("Open Cursor!")
    }
}

abstract class Application {

    // Factory Method
    abstract createEditor(): Editor 

    // Templeate method
    newEditor() {
        const doc = this.createEditor();
        doc.open();
    }
}

// Concrete Creators
class TextEditorApp extends Application {
    createEditor(): Editor {
        return new VsCode();
    }
}

class SheetApp extends Application {
    createEditor(): Editor {
        return new Cursor();
    }
}

// use case
const app1: Application = new TextEditorApp();
app1.createEditor();
app1.newEditor();

const app2: Application = new SheetApp();
app2.createEditor();
app2.newEditor();


// Key point ของ Factory Method

// มี method ที่เอาไว้สร้าง object อยู่ใน base class (เช่น createEditor())

// แต่ subclass เป็นคนระบุ ว่าจะ return อะไร (TextEditorApp vs SheetApp)

// client ใช้ผ่าน base class (Application) ได้เลย ไม่ต้องรู้ว่าข้างหลังคือ TextEditorApp หรือ SheetApp