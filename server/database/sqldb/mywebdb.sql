use mywebdb

GO

drop table if exists Accounts;

drop table if exists Tokens;

drop table if exists Tables;

create table Tables (
    ID  INT NOT NULL IDENTITY UNIQUE,
    dep VARCHAR(5) NOT NULL CHECK (dep like 'acctk' or dep like'token')
);

create table Tokens (
    token   NVARCHAR(40) PRIMARY KEY,
    tabid   INT NOT NULL REFERENCES Tables(ID),
    cdate   DATETIME NOT NULL
);

create table Accounts (
    ID      INT NOT NULL IDENTITY UNIQUE,
    token   NVARCHAR(40) NOT NULL UNIQUE REFERENCES Tokens(token),
    fname   NVARCHAR(50) NOT NULL,
    uname   NVARCHAR(20) NOT NULL UNIQUE,
    passw   NVARCHAR(20) NOT NULL
);

GO

insert into Tables (dep) values ('token');
insert into Tables (dep) values ('acctk');


GO

create or alter procedure AddToken
@token NVARCHAR(40),
@tabid INT
as 
begin
    if exists (
        select token 
        from Tokens
        where token = @token
    ) 
        return;
    else
        insert into Tokens
        values (@token, @tabid, GETDATE());
end
GO

create or alter procedure RemoveToken
@token NVARCHAR(40),
@tabid INT,
@deepremove BIT = 0
as
    if @tabid = 2 and @deepremove = 0
        return
    else 
        delete from Tokens
        where tabid = @tabid and token LIKE @token 

GO

create or alter function CheckToken
(
    @token NVARCHAR(40),
    @tabid INT
)
returns table
as 
    return (
        select * 
        from Tokens
        where tabid = @tabid and token like @token
    )  

GO

create or alter procedure AddAccount
@token NVARCHAR(40),
@fname NVARCHAR(50),
@uname NVARCHAR(20),
@passw NVARCHAR(20)
as
begin 
    begin tran addaccount;
    begin try
        if exists (
            select *
            from CheckToken(@token, 2)
        )
            THROW 51000, 'token already exists.', 1;

        exec AddToken @token, 2;

        if not exists (
            select *
            from CheckToken(@token, 2)
        )
            THROW 51000, 'add token fails.', 1;

        insert into Accounts(token, fname, uname, passw)
        values(@token, @fname, @uname, @passw);

        commit tran addaccount;
    end try 
    begin catch
        rollback; 
    end catch
end

GO

create or alter procedure AddOOAccount
@token NVARCHAR(40),
@fname NVARCHAR(50),
@uname NVARCHAR(20),
@passw NVARCHAR(20)
as
begin 
    begin tran addooaccount;

    if exists (
        select *
        from CheckToken(@token, 2)
    )
    begin 
        update Accounts
        set fname = @fname, uname = @uname, passw = @passw
        where token = @token;
    end 
    else 
    begin 
        exec AddAccount @token, @fname, @uname, @passw;
    end 

    commit tran addooaccount;
end

GO

create or alter procedure RemoveAccount
(
    @token NVARCHAR(40)
)
as
begin 
    if not exists (
        select *
        from CheckToken(@token, 2)
    )
        return
    
    begin tran removeaccount;
    begin try     
        delete from Accounts
        where token like @token;

        exec RemoveToken @token, 2, 1

        commit tran removeaccount;
    end try 
    begin catch 
        rollback;
    end catch 
end

GO

create or alter procedure UpdateAccount
@token NVARCHAR(40),
@fname NVARCHAR(50),
@uname NVARCHAR(20),
@passw NVARCHAR(20)
as 
begin
    if not exists (
        select *
        from CheckToken(@token, 2)
    )
        return;
    
    begin tran updateaccount;
    begin try
        update Accounts
        set fname = @fname, uname = @uname, passw = @passw
        where token = @token

        if not exists (
            select *
            from Accounts
            where token = @token and fname = @fname and uname = @uname and passw = @passw
        )
            THROW 51000, 'token already exists.', 1;
        
        commit tran updateaccount;
    end try 
    begin catch 
        rollback; 
    end catch
end 

GO

create or alter function CheckAccount
(
    @uname NVARCHAR(20),
    @passw NVARCHAR(20)
)
returns table 
as 
    return (
        select [Accounts].[ID],
        [Accounts].[token],
        [Accounts].[fname],
        [Accounts].[uname],
        [Accounts].[passw],
        [Tokens].[cdate]
        from Accounts
        inner join Tokens
        on Accounts.token = Tokens.token
        where uname = @uname and passw = @passw
    )
    
GO

create or alter function CheckTkAccount
(
    @token NVARCHAR(40)
)
returns table 
as 
    return (
        select [Accounts].[ID],
        [Accounts].[token],
        [Accounts].[fname],
        [Accounts].[uname],
        [Accounts].[passw],
        [Tokens].[cdate]
        from Accounts
        inner join Tokens
        on Accounts.token = Tokens.token
        where Accounts.token = @token
    )
    
GO

exec AddToken '0000000000000000', 1;

GO