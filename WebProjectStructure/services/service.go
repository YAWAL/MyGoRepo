package services

/*
https://habrahabr.ru/post/331456/?utm_source=fb&utm_medium=social&utm_campaign=kak-my-homyaka-yablokami-kormili-ili-effek

Ситуация с сервисами во многом идентична, они так же начинаются с интерфейса, структуры и конструктора с последующим
набором методов. Акцент хочется сделать на одной детали — это принцип работы с базой.

Конструктор сервиса должен принимать в себя набор репозиториев, с которыми он будет работать, и фабрику транзакций:

func NewOrder(repo repositories.Order, txFactory TransactionFactory) Order {
    return &order { repo: repo, txFactory: txFactory }
}

Фабрика транзакций — это просто класс, генерирующий транзакции, здесь ничего сложного:

type TransactionFactory interface {
    BeginNewTransaction() Transaction
}

Полный код фабрики для gorm
type TransactionFactory interface {
    BeginNewTransaction() Transaction
}

type transactionFactory struct {
    db *gorm.DB
}

func NewTransactionFactory(db *gorm.DB) TransactionFactory {
    return &transactionFactory{db: db}
}

func (t transactionFactory)BeginNewTransaction() Transaction {
    tx := new(transaction)
    tx.db = t.db
    tx.Begin()
    return tx
}

А вот на самих транзакциях остановиться стоит. Начнем с того, что это вообще такое. Транзакция представляет из себя
тот же интерфейс с реализацией, который содержит методы для старта транзакции, завершения, отката и доступа к
реализации движка уровнем ниже:

type Transaction interface {
    Begin()
    Commit()
    Rollback()
    DataSource() interface{}
}

Полный код транзакции для gorm
type Transaction interface {
    Begin()
    Commit()
    Rollback()
    DataSource() interface{}
}

type transaction struct {
    Transaction
    db *gorm.DB
    tx *gorm.DB
}

func (t *transaction)Begin() {
    t.tx = t.db.Begin()
}

func (t *transaction)Commit() {
    t.tx.Commit()
}

func (t *transaction)Rollback() {
    t.tx.Rollback()
}

func (t *transaction)DataSource() interface{} {
    return t.tx
}

Если с begin, commit, rollback все должно быть понятно, то Datasource — это просто костыль для доступа к
низкоуровневой реализации, потому что работа с любой БД в Go устроена так, что транзакция является просто копией
акссессора к базе со своими измененными настройками. Он нам понадобится позже при работе в репозиториях.

Собственно, вот и пример работы с транзакциями в методе сервиса:

func (o order)PlaceOrder(m *model.Order)  {
    tx := o.txFactory.BeginNewTransaction()
    defer tx.Commit()
    o.repo.Insert(tx, m)
}

Начали транзакцию, выполнили доступ к базе, закоммитили или откатили, как больше нравится.

Конечно, все преимущество транзакций особенно раскрывается при нескольких операциях, но и даже если у вас всего одна,
как в примере, хуже от этого не будет.

Экспертам
Знаю, что нет управления уровнями изоляции.
Если нашли еще какие косяки — пишите в комментах.

В качестве дополнительного совета юниорам, хочу сказать, что транзакция должна быть открыта минимально возможное время.
Постарайтесь подготовить все данные так, чтобы на период между begin и commit приходилось минимальное количество логики
и вычислений.
Бывает, что транзакцию открывают и идут курить, отправляя, например запрос в гугл. А потом удивляются, почему это с
дедлоком зафакапилось все.

Интересный факт
Во многих современных базах данных, deadlock определяется максимально просто: по таймауту. При большой нагрузке
сканировать ресурсы на предмет определения блокировки — дорого. Поэтому часто вместо этого используется обычный таймаут.
Например, в mysql. Если не знать эту особенность, то можно подарить себе чудеснейшие часы веселой отладки.

 */
