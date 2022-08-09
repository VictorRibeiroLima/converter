# Converter

Golang Simple type converter between structs.

This package tries to match name and type of the properties on a struct to make the conversion

## Usage

Let's say you have a DTO that needs to be converted to a Model.

```go
package main


import (
	"fmt"

	"github.com/VictorRibeiroLima/converter"
)

type User struct {
	ID       uint
	UserName string
	Password string
	IsAdmin  bool
}

type CreateUserDto struct {
	UserName string
	Password string
}
```

If the DTO and the Model are not a perfect match you can't do the casting.

```go
func main() {
	var user User
	dto := CreateUserDto{
		UserName: "LukeSlywalker",
		Password: "benKenoby1977",
	}

	user = (User)(dto)//cannot convert dto (variable of type CreateUserDto) to User
	fmt.Println(user)
}
```
Well that's why this package was made,to make those simple conversions

```go
func main() {
	var user User
	dto := CreateUserDto{
		UserName: "LukeSlywalker",
		Password: "benKenoby1977",
	}

	converter.Convert(&user, dto)

	fmt.Println(user) //{0 LukeSlywalker benKenoby1977 false}

  /*
    User{
      ID: 0,
      UserName: "LukeSlywalker",
      Password: "benKenoby1977",
      IsAdmin: false
    }
  */
}
```
## Advanced Usage

Let's dive a bit deeper on the use cases.

What if you have a nested struct inside your Model And DTO,and let's throw some pointer in there too (GORM use this to make NOT NULL fields on databases)

```go
package main

import (
	"fmt"

	"github.com/VictorRibeiroLima/converter"
)

type Address struct {
	ID      uint
	Country string
	City    string
	Street  string
}

type User struct {
	ID       uint
	UserName *string
	Password *string
	IsAdmin  bool
	Address  Address
}

type AddressDto struct {
	Country string
	City    string
	Street  string
}

type CreateUserDto struct {
	UserName string
	Password string
	Address  AddressDto
}

func main() {
	var user User
	address := AddressDto{
		Country: "Jundland Wastes",
		City:    "Great Chott salt flat",
		Street:  "Lars homestead",
	}
	dto := CreateUserDto{
		UserName: "LukeSlywalker",
		Password: "benKenoby1977",
		Address:  address,
	}

	converter.Convert(&user, dto)

	fmt.Println(user)

  /*
    User{
      ID: 0,
      UserName: 0xc000010250,
      Password: 0xc000010260,
      IsAdmin: false,
      Address: {
        ID 0,
        Country: "Jundland Wastes",
	City:    "Great Chott salt flat",
	Street:  "Lars homestead",
      }
    }
  */
}
```
```diff
- (user.UserName user.Password) will not be pointers to (dto.UserName dto.Password)
```

You can do too:
- array to array
- slice to slice
- poiter to value
- value to poiter
- pointer to pointer
- slice of struct to slice of another struct
- slice of pointer of struct to slice of value of another struct

You get the idea
