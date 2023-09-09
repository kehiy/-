const file = Bun.file("./test.txt")

console.log("type:")
console.log(file.type)
console.log("size before write:")
console.log(file.size)

Bun.write("./test.txt", "Hello, world! #!%$@%$&")

console.log("size after write:")
console.log(file.size)

// console.log("json:")
// console.log(await file.json()) // this option pareses json!
console.log("text:")
console.log(await file.text())
