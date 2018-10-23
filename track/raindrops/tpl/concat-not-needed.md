- You could check your usage of `+=` in regards to where it is really needed. If the target variable is 
guaranteed to be empty anyway simple `=` does the job and is faster.