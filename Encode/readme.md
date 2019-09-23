## TLV

- `tag`    unique identify
- `length` length of value
- `value`  data self

---

` type ,length = 2 byte,4 byte ; value is depend on length`


```bash
    ## socket func
    htonl() #Host to Network Long
    ntohl() #Network to Host Long
    htons() #Host to Network Short
    ntohs() #Network to Host Short   
```

1. 网络字节顺序NBO（Network Byte Order）\
   从高到低的顺序存储
    
2. 主机字节顺序HBO (Host Byte Order）\
   与CPU设计有关,数据的顺序由CPU决定
   
   ```bash
   Intelx86 : short 0x1234 -> 34 12, int 0x12345678 -> 78 56 34 12 little endian
   IBM power PC : short 0x1234 -> 12 34, int 0x12345678 -> 12 34 56 78 big endian

   ```
   
   
   inter 
   ```bash
   # winsock2.h
     int main()
       {
           printf("%d \n",htons(16)); # 0x0010
           return 0;
       } # printf 4096
   ```
  
  > encode order : tag -> length -> value \
  
  1. BER 编码 （base encoding rule）
  
        1.> 长度确定的编码方式：
        Identifier octets,
        Length octets,
        Contents octets \
        2.> 长度部确定的编码方式:
        Identifier octets,
        Length octets (0x80), 
        Contents octets,
        End-of-contents octets(0x00 00)
  
  2. DER 编码   (distinguished encoding rules)
  
        仅长度确定的编码
  
  ---
  
  `Identifier octets ：由3部分组成 Class,P/C,Tag number` 
  
         第一个字节的高2位为Class,接下来一位为P/C,其他位表示Tag number
  
  `Class: 有4中类型Universal(00),Application(01),Context-specific(10),Private(11)`
  
         P/C : 位如果为1则表示是Constructed的,为0表示是Primitive
         
         if 0<=Tag number<=30 { 
               Identifier octets  1 bytes 
         } else { 
               第一个字节的后5位前为1 
               后面第一个最高位为0的字节,该字节就是Identifier octets的最后一个字节 
         } 
         第二个字节到最后一个字节去掉最高位的值拼起来就是 Tag number的值
         
         长度确定的编码方式的Length octets有两种方法编码长度,一种是只用一个字节表示长度,其最高位为0,后7位表示长度值,显然这样只能表示0-127
         
         另一种是第一个字节的最高位为1,其他位表示后面还有多少个字节属于Length octets,后面的那些字节组成的就是长度值,
         长度值表示的是Contents octets所占的字节数 
         DER要求如果长度为0-127则要使用第一种方式,如果大于127则使用后一种方式
   
  `Length: 4 bytes`
  
        如果第一个字节的最高位b8为0, b7~b1的值就是value域的长度 
        如果b8为1,b7~b1的值指示了下面有几个子字节,下面子字节的值就是value域的长度
   

  
  
  