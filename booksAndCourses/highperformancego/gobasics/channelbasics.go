package main

import fmt

func main(){
	buffchannel:=make(chan string, 2)
	buffchannel<-"foo"
	buffchannel<-"bar"

	fmt.Println("size of channel buffchannel is ", len(buffchannel))

	fmt.Println(<-buffchannel)
	fmt.Println(<-buffchannel)

	//length of channel is 0 because both elements removed from the channel
	fmt.Println("channel lenght after popping strings from it ",len(buffchannel))

	buffchannel<-"baz"

	out:=<-buffchannel
	fmt.Println(out)

close(buffchannel)
}