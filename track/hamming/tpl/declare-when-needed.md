- you are declaring `%s` before you check if the length is equal. In case the length is not equal it is 
declared in vain. Best declare variables when they are needed, not before.