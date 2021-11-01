package revrpc

type ReverseServer struct{}

func (server *ReverseServer) Add(nums [2]int, reply *int) error {
	*reply = nums[0] + nums[1]
	return nil
}
