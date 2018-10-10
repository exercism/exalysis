- every value type can be formatted with the place-holder `%%v`. One of Go's strengths is the type system.
You have the advantages of Go's type system in formatting as well if you use the correct placeholder for 
each type. A strings placeholder is `%%s`. The type specific placeholders also offer more options for 
formatting that type. For example the placeholder for floats `%%f` can also specify the precision: `%%.2f`.