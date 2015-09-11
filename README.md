# 七牛云存储二面 Open Question

https://github.com/wangkechun/qiniu_1

### 异进制加法：输入两个数以及其对应的进制，输出两数之合及其进制。



#### 比如：
输入： 12 3 11 5 对应，
输入为 3 进制的 12 和 5 进制的 11

输入： 11 10，
输出为 10 进制的 11

#### 要求（以下要求区分先后顺序，要优先满足标号小的要求）

- 1 整数加法
- 2 小数加法，不能损失精度
- 3 超过 36 进制/62 进制，数的输出
- 4 超过 36 进制/62 进制，数的输入
-（ 语言不限，使用最熟悉的语言。 ）

#### 思路：
>不同进制小数加法，要求不能损失精度

>可以转换成它们进制的最小公倍数再运算就可以规避精度问题。

>小数的进制转换处理较为麻烦，具体细节就不解释了。

>如果输入是整数的话，结果可以转换成任意进制。

>输入超过36进制的数字，采用[1,2.3,4]来输入,

>否则可以使用10个数字加26个字母来表示(不区分大小写).

>输出逻辑同输入。

#### 注意点：
输出的时候去掉首尾的0， 整数输出不包含".0"
进位处理

#### 运行：
```
git clone https://git.oschina.net/wkc/qiniu_1.git
make install
make test
```


生成测试数据可以使用mathematica等工具:

http://www.wolframalpha.com/

http://reference.wolfram.com/language/tutorial/DigitsInNumbers.zh.html

#### godoc
```
PACKAGE DOCUMENTATION

package bignum
    bignum 七牛云存储二面 Open Question 异进制加法：输入两个数以及其对应的进制，输出两数之合及其进制。

TYPES

type BigNum struct {
    // contains filtered or unexported fields
}
    BigNum 大数存储

func New(s string, base int) (v BigNum)
    New returns a new BigNum

func (x *BigNum) Add(y *BigNum) (z BigNum)
    Add 大数相加

func (x *BigNum) ChangeBase(newBase int) error
    ChangeBase 修改数字的进制

func (x *BigNum) Format()
    Format 格式化大数， 去掉多余的0

func (x *BigNum) Input(s string, base int) error
    Input 输入一个大数

func (x *BigNum) RawString() (r string)
    RawString 转换成 [1,2,3] 类似的形式

func (x *BigNum) String() (r string)
```
