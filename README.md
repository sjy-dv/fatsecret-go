## Documentation

fatsecret's api documentation is not friendly.
This is a document organized by creating an API that can be used in golang.

## Explain

```bash
// The api will be updated continuously.

*** currently applied api ***
- Authorization Token // Includes ongoing token reissuance and maintenance without worrying about token expiration.

- food.get.v2
- food_categories.get
- food_sub_categories.get
- foods.search
```

## Sample Code

```go

retval, err := fatsecret_go.FoodCategoriesGet()
if err != nil {
    fmt.Println(err)
}
fmt.Println(retval)
```
