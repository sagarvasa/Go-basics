# Control flow - Summary

There are 2 fundamental statements
1. If statement
2. Switch statement

# if statement
1. Similar to if condition of most of the language except condition needs to be boolean only and can not be alias for any other type

2. Curly braces around if block is mandatory

3. types are -> if , if-else , if-else if-else


# switch statement
1. There are multiple type of switch statements
    1. Switch with tags
    2. Switch without tags
    3. Type Switch

2. In switch statement break is implicit and fallthrough is explicit (unlike most of other programming lang)

3. case can not be overridden. it will throw compiler error if we try to define multiple same case at the same time

4. In case of switch without tags, case should present boolean condition. We can also override case in case of switch without tags

5. we can use type switch along with interface to identify types and perform corresponding action