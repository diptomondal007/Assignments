package adapter_practice

type Client struct {

}

func(c *Client) InsertSquareUsbIntoComputer(com Computer){
	com.InsertInSquarePort()
}
