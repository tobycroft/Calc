# Calc from TuuzGoWeb

这是TuuzGoWeb中的Calc模块，用于数据类型转换，使用简单方便

# 使用方法

- go.mod 中添加

```bash
require github.com/tobycroft/Calc v1.0.0
```

- go get

```bash
go get -u github.com/tobycroft/Calc
```

# 功能

- 类型转换
    - 任意转换(Any2)
        - Calc.Any2....
            - ret:=Calc.Any2String(anyinterface)
            - ret:=Calc.Any2Float64(anyinterface)
            - ret:=Calc.Any2Int64(anyinterface)
        - Calc.Any2..._2
            - ret,err:=Calc.Any2Int64_2("string_number")
            - ret,err:=Calc.Any2Float64_2("string_number")
    - 精密计算(Bc_)
        - 任意转换成精密数字
            - decimal_val:=Calc.todecimal(decimal_number1)
        - 加法
            - decimal_val:=Calc.Bc_add(decimal_number1,decimal_number2)
            - decimal_val:=Calc.Bc_sum(decimal_number1,decimal_number2)
        - 减法
            - decimal_val:=Calc.Bc_min(decimal_number1,decimal_number2)
        - 乘法
            - decimal_val:=Calc.Bc_mul(decimal_number1,decimal_number2)
        - 除法
            - decimal_val:=Calc.Bc_div(decimal_number1,decimal_number2)
            - decimal_val:=Calc.Bc_div_round(decimal_number1,decimal_number2,round_number)
        - 精确小数点
            - decimal_val:=Calc.Bc_round(decimal_number1,decimal_number2)
        - 取余
            - decimal_val:=Calc.Bc_mod(decimal_number1,decimal_number2)
        - 开平方/开方
            - decimal_val:=Calc.Bc_pow(decimal_number1,decimal_number2)