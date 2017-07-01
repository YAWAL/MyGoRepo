package repositories

/*
https://habrahabr.ru/post/331456/?utm_source=fb&utm_medium=social&utm_campaign=kak-my-homyaka-yablokami-kormili-ili-effek

Тоже самое: интерфейс, структура, конструктор, который, как правило, уже без параметров.
Просто приведу пример операции Insert, которую мы вызывали в коде сервиса:

func (order)Insert(tx Transaction, m *model.Order) {
    db := tx.DataSource().(*gorm.DB)
    query := "insert into orders (shop_id) values (?) returning id"
    db.Raw(query, m.Shop.ID).Scan(m)
}

Получили из транзакции низкоуровневый модификатор доступа, составили запрос, выполнили его. Готово.

Всего этого вполне должно хватить, чтобы не угробить архитектуру. По крайней мере слишком быстро.
 */
