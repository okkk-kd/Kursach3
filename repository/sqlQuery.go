package repository

const (
	queryCreateUser = `
	insert into users(username, password, role_id) values($1, $2, 1)
	`
	queryUpdatePassword = `
	update users
	set password = $1
	where id = $2
	  and password = $3
	returning true;
	`
	queryAuthorizationInsertSessionKey = `
	insert into sessions(token, active, user_id)
	values ($1, $2, $3);
	`
	queryUpdateSessionKey = `
	update sessions
	set authed = false
	where user_id = $1
	  and authed = true;
	`
	queryCheckUserParams = `
	select case when id = $1 and password = $2 then true else false end as authed
	from users
	where id = $1;
	`
	queryWithdraw = `
	insert into orders(from_wallet_id, to_wallet_id, quontity, currency_id, card_id) values($1,$2,$3,$4,$5);
	`
	queryCreateCard = `
	insert into cards("number", validated) values($1, false);
	insert into bank_account(card_id, user_id) values($1, $2);
	`
	queryGetCurrencies = `
	select id, name from currencies c;
	`
	queryGetCards = `
	select c."number" , c.id, c.wallets, c.validated  from cards c  where c.id = $1 and c."blocked" = false;
	`
	queryGetBlockedCards = `
	select c."number" , c.id, c.wallets, c.validated  from cards c  where c.id = $1 and c."blocked" = true;
	`
	queryGetUserInfo = `
	select u.username, u.active, u.role_id from users u where u.id = $1;
	`
	queryGetOrders = `
	select o.from_wallet_id , o.to_wallet_id , o.quontity , o.currency_id  from orders o where o.user_id = $1;
	`
	queryGetOrderInfo = `
	select o.from_wallet_id , o.to_wallet_id , o.quontity , o.currency_id  from orders o where o.user_id = $1 and o.id = $2;
	`
	queryGetCardOrders = `
	select o.from_wallet_id , o.to_wallet_id , o.quontity , o.currency_id  from orders o where o.user_id = $1 and o.card_id  = $2;
	`
)
