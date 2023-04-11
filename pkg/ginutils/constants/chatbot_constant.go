// Không check long line linter cho message
// nolint: lll
package constants

const (
	ShopFPTID               = "888888"
	ShopLogo                = "https://media3.scdn.vn/img4/2022/05_18/a3xuwrVqvkn999LRjFFu.png"
	ShopName                = "Trợ lý Farm"
	ChatbotExtraType        = "orders"
	AttachmentType          = "template"
	ChatBotTemplate         = "button"
	ButtonType              = "web_url"
	ActionSetPackageTitle   = "Mở Lì xì"
	ActionPickPackageTitle  = "Mua hàng ngay"
	ActionToFarmWalletTitle = "Kiểm tra ví"
)

// Nội dung chatbot cho arlert
const (
	TopupFailedMsg = "```[ALERT] Đơn hàng %s hoàn tiền lỗi, kiểm tra ngay.! ErrorMsg: %v ``` Splunk Link:%s"
	SplunkErrorSTG = "https://bit.ly/3WyMF2h"
	SplunkError    = "https://bit.ly/3Whr3aW"
)

// Nội dung bắn chatbot
const (
	ChatBotAddRedPacketMsg  = `<b>%s %s nhận được 1 Lì xì may mắn</b><br><br>ết đến xuân về, Sendo Farm thân tặng %s 1 bao Lì xì may mắn. Mở ngay để xem lộc đầu năm của mình %s nhé.`
	ChatBotBuyFirstOrderMsg = `<b>Mua 1 đơn hàng để tích lũy %sđ</b><br><br>Chúc mừng %s %s đã tích được %sđ vào Lì xì. Mua ngay 1 đơn hàng trước 23h59 ngày %s để tích thêm %sđ %s nhé.`
	ChatBotBuyNextOrderMsg  = `<b>Mua 1 đơn hàng để tích lũy %sđ</b><br><br>%s %s ơi, hiện tại %s đã tích lũy được %sđ vào Lì xì. Tiếp tục mua 1 đơn hàng trước 23h59 ngày %s để tích thêm %sđ %s nhé.<br><br>Khi tích đủ %sđ, %s sẽ được nhận số tiền này vào Ví Farm khuyến mãi đấy.`
	ChatBotBuyLastOrderMsg  = `<b>Mua 1 đơn hàng để nhận đủ Lì xì %sđ</b><br><br>%s %s ơi, chỉ còn <b>1</b> đơn hàng từ %sđ nữa là mình sẽ tích lũy đủ Lì xì.<br><br>Tiếp tục mua vào ngày mai (%s) để nhận đủ Lì xì %sđ vào Ví Farm khuyến mãi %s nhé.`
	ChatBotCanClaimMsg      = `<b>Nhận %sđ Lì xì về Ví Farm khuyến mãi</b><br><br>Chúc mừng %s %s đã tích lũy đủ %sđ tiền Lì xì. Nhận tiền ngay về Ví Farm khuyến mãi để sử dụng %s nhé.<br><br>Sau hh:mm ngày %s, %s sẽ không thể bấm nhận tiền được nữa, nên mình tranh thủ nhận nhé.`
	ChatBotPickPackageMsg   = `<strong>Mua 1 đơn hàng từ %sđ để tích lũy %sđ</strong><br><br>Chúc mừng %s %s đã tích được <strong>%sđ</strong> vào Lì xì. Để tích thêm <strong>%sđ</strong> nữa, mua ngay 1 đơn hàng từ %sđ <strong>trước 23h59 ngày %s</strong> %s nhé.<br><br>Khi tích đủ %sđ, %s sẽ được nhận số tiền này vào Ví Farm khuyến mãi đấy.`
	ChatBotPickPackageDesc  = `<strong>Lì xì tích lũy</strong>`
	ChatBotClaimMoneyMsg    = `<strong>Lì xì %sđ đã cộng vào Ví Farm khuyến mãi </strong><<br><br>>Tuyệt quá, %sđ tiền lì xì đã được cộng vào Ví Farm khuyến mãi của %s.<br><br>Số tiền này có hạn dùng đến hết ngày <strong>%s</strong>, nên %s tranh thủ dùng mua rau củ, trái cây, thực phẩm tươi sống, thực phẩm chế biến sẵn nhé.`
)
