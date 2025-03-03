# **Golang Card Validator**

<div align="center">
    <a href="#">  <img height="150" src="https://raw.githubusercontent.com/betandr/gophers/master/Gopher.png" alt="Gopher" /></a>
</div><br>

## **Description**

The **Card Validator** is a package for validating credit card numbers using the **Luhn algorithm**. In addition to validation, the package identifies the card brand based on common prefixes.

## **Usage**

### **Basic Example**

```go
package main

import (
    "fmt"
    "github.com/JrSchmidtt/card-validator"
)

func main() {
    validator := card_validator.NewCardValidator()
    isValid, brand := validator.Validate("5158820011790888")
    
    fmt.Println("Valid card:", isValid)
    fmt.Println("Brand:", brand)
}
```

### **Example with Custom Prefixes**

You can customize the card prefixes by passing a map of prefixes to the `NewCardValidator` function. Here's an example:

```go
package main

import (
    "fmt"
    "github.com/JrSchmidtt/card-validator"
)

func main() {
    customPrefixes := map[string]string{
        "6062": "hipercard",
    }

    validator := card_validator.NewCardValidator(customPrefixes)
    isValid, brand := validator.Validate("6062-8268-1892-9606")
    
    fmt.Println("Valid card:", isValid) // true
    fmt.Println("Brand:", brand) // hipercard

    isValid, brand = validator.Validate("2212345678901234")
}
```


## **Contributing**

1. [Fork the repository](https://github.com/JrSchmidtt/card-validator/fork)
2. Clone your fork:  
   ```sh
   git clone https://github.com/JrSchmidtt/card-validator.git
   ```
3. Create a new branch for your feature:  
   ```sh
   git checkout -b my-new-feature
   ```
4. Make your changes and commit:  
   ```sh
   git commit -am 'Add new feature'
   ```
5. Push to the remote branch:  
   ```sh
   git push origin my-new-feature
   ```
6. Open a **Pull Request** 🚀

## **Author**
Maintained by [Junior Schmidt](https://github.com/JrSchmidtt).

