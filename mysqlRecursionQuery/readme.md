## Mysql Recursion 

### 场景

- 微博评论
- 商品目录
- ...


### 准备工作

[创建 comment 表]() \
[删除 comment 表]() \
[插入 测试数据 & 添加 父节点索引]()

```
            |---->LETTER---->letter(list)
ALL--root-->|
            |---->NICKS---->NICK---->nick(list)
            
ALL root id = 1

LETTER parent_id = 1, id = 3
NICKS parent_id = 1, id = 2
NICK parent_id = 2, id = 4

letter parent_id = 3
nick parent_id = 4
```


### 数据结构

    一般情况下，我们需要保证每条数据存入时： 
    其中 `id` 表示 唯一ID, parent_id 表示父节点的 `id`
    
```
| id | other_column ... | parent_id | level (option) |
```
    
### 递归函数

- prepare work

```sql
set global log_bin_trust_function_creators=TRUE;
```

#### up recursion

> `函数名称：` `UpFetchCommentTree`

```sql
drop function if exists UpFetchCommentTree;
create function UpFetchCommentTree(rootID INT)
returns varchar(4096)
begin
  declare Parent varchar(4096);
  declare Children varchar(4096);
set Parent='$';
set Children = cast(rootID as char);
set Parent = CONCAT(Parent,',',Children);

select parent_id into Children
  from comment where id = Children;

while Children <> 0 do
  set Parent = concat(Parent,',',Children);
  select parent_id into Children
    from comment where id = Children;
end while;
return Parent;
end;

```

#### down recursion

> `函数名称：` `DownFetchCommentTree

```sql
drop function if exists DownFetchCommentTree;
create function DownFetchCommentTree(rootID varchar(32))
returns varchar(4096)
begin
  declare Parent varchar(4096);
  declare Children varchar(4096);
set Parent='$';
set Children=cast(rootID as char);

while Children is not null do
  set Parent=concat(Parent,',',Children);
  select group_concat(id) into Children
    from comment where find_in_set(parent_id,Children)>0;
end while;
return Parent;
end;

```       

- check function

```sql
show function status  where Name = 'UpFetchCommentTree';
show function status  where Name = 'DownFetchCommentTree';
```

- up recursion 

    通过节点向上查询 `target_id`

```sql
select * from comment where find_in_set(id,UpFetchCommentTree(3));
```
    
- down recursion

    通过节点向下查询 `target_id`
        
```sql
select * from comment where find_in_set(id,DownFetchCommentTree(3));
```
