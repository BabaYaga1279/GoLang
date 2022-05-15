use mywebdb

GO

exec AddOOAccount '0000000000000001', 'minh bui', 'bqm2709', 'Quangminh270901';


select * from CheckAccount('bqm2709', 'Quangminh270901')