package common

type Message struct {
	ProductId int
	UserId    int
}

// 创建结构体
func NewMessage(userId, productId int) *Message {
	return &Message{ProductId: productId, UserId: userId}
}
