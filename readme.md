# OneIota go challenge
## Structure
In terms of structure, I went for a 3 layer application, where the `main` function serves as the controller layer. 

The service layer retrieves data in the correct format from the repo, applies any business logic, and returns the data to the controller. 

The repository layer extracts data from the CSV file, and maps it to a product struct that match the specified structure.

Models and any logic surrounding them are in their own package. Some of the models will implement the Sort interface in order to define custom sorting criteria.

I have also created a util package for code that doesn't need to be directly in the repository layer for code clarity.

I made sure to include some basic error handling in the endpoints, and `panic`ked out of the application if the data in the CSV file was invalid, otherwise any errors were consumed and logged.

To add new size options, simply add another case to the `switch case` in `repo/product_repo.go` and define a new size map with a custom order in `util/size_order_util.go`

### Pitfalls
The API works under the assumption that all the data is sorted by it's PLU. In this case it is, but in future iterations, I would load all data at once using `reader.ReadAll(file)` and ensure to sort the data before processing.

## Thought processes
When iterating over the data, check to see if the PLU matches the PLU from the previous line, if so add the size to the existing product, if not, make a new product and add the size.

#### Structure should be

```
product {
    PLU: string,
    name: string,
    sizes: {
        SKU: int,
        size: string
    }
}
```

#### CVS data order
```
PLU:      record[1],
SKU:      record[0],
Name:     record[2],
Size:     record[3],
SizeSort: record[4],
```

### Main loop over data
To ensure the data matches the desired structure, we'll need to keep track of the sizes of any products with matching IDs while iterating over the records.
Once we encounter a non-matching ID, we create a new product with the sorted array of sizes, and append it to the list of products, and finally reset the sizes array with the current record's value.

### Issue: Quotes
Without `reader.LazyQuotes` set to true, we get errors when parsing the CSV data.
Once set to true, quotes are escaped and we end up with `\"` chars surrounding each field.
This causes issues when trying to map them to the size map as the strings aren't equal.
To overcome this issue, I will trim each entry of escaped quotes as well as spaces:

```
for i, field := range record {
    cleaned := strings.TrimSpace(field)
    cleaned = strings.Trim(cleaned, "\"")

    record[i] = cleaned
}
```

in `getRecord()`, `repo/product_repo.go`

### Issue: Last product is being ignored
Probable cause: Once the last line is reached, the loop breaks and the product is not added to the slice.
Current state of the code:

```
for i := 0; ; i++ {
    previousPLU := previousRecord[1]
    currentPLU := currentRecord[1]

    if previousPLU == currentPLU || i == 0 {
        sizes = append(sizes, getSize(currentRecord))
    } else {
        sortSizes(sizes, previousRecord[4])
        product := model.Product{
            PLU:   previousPLU,
            Name:  previousRecord[2],
            Sizes: sizes,
        }
        products = append(products, &product)
        sizes = []model.Sizes{getSize(currentRecord)}
    }

    previousRecord = currentRecord
    currentRecord = getRecord(reader)
    if currentRecord == nil {
        break
    }
}
```

#### Potential fixes:
Keep track of whether the last item has been reached with a boolean and continue iterating.
I don't like the above fix as it adds more branches to the code and makes it a bit less readible.

#### Actual fix:
I'm going to break out the product creation logic to a helper function and once the last item in the file has been reached, I'll call the helper method and append the last product to the slice before breaking out of the loop.

#### Fix applied:
After this fix, the code is more concise as some logic has been broken out of the loop makking it more readible, and the code is able to be reused to add the final product to the slice once the end of file is reached.

### Issue: Child size orders are incorrect
Because children sizes are followed up by brackets, they are lexicographically greater than sizes without that string, so in the custom `Less` sorting method in `model/sizes.go`, we have to check whether that string is present and make sure that those values are counted as less than values without the `(Child)` postfix present.

### Issue: 10 is coming after 9 in sorting
Beacuse I was trying to sort using the `string` value of the sizes, 10 was being evaluated as smaller than 9 due to the 1 at the begining of the string. To avoid this issue, the values are converted to floats for comparison in the `Less` method of `model/sizes.go`
