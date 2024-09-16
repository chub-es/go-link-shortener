create or replace function generateShortUrl(len integer) returns text as
$$
declare
	chars text[] := '{0,1,2,3,4,5,6,7,8,9,A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z,a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z}';
	s text := '';
	i integer := 0;
begin
	-- control
	if len < 0 then
    	raise exception 'Указанная длина строки не может быть меньше 0.';
  	end if;

	-- generate
	for i in 1..len loop
		s := s || chars[1+random()*(array_length(chars, 1)-1)];
	end loop;

	-- return
	return s;
end;
$$ language plpgsql;

create table of not exists links(
    id serial primary key,
    created_at timestamp default now(),
    original_url varchar(2000) not null,
    short_url varchar(50) not null unique default generateshorturl(6),
    showned integer default 0
);