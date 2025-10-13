import "container/list"

type Twitter struct {
	totalTweets int
	users       map[int]*user
}

type user struct {
	last10Tweets *list.List
	follows      map[int]bool
}

func newUser() *user {
	return &user{
		last10Tweets: list.New(),
		follows:      make(map[int]bool),
	}
}

func Constructor() Twitter {
	return Twitter{
		totalTweets: 0,
		users:       make(map[int]*user),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.createUserIfNotExists(userId)

	user := this.users[userId]

	user.last10Tweets.PushFront([2]int{this.totalTweets, tweetId})
	this.totalTweets++

	if user.last10Tweets.Len() == 11 {
		user.last10Tweets.Remove(user.last10Tweets.Back())
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	this.createUserIfNotExists(userId)

	user := this.users[userId]
	heap := newMaxHeap()

	heap.PushTweets(user.last10Tweets)

	for follower, _ := range user.follows {
		heap.PushTweets(this.users[follower].last10Tweets)
	}

	return heap.ExtractTweets()
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}

	this.createUserIfNotExists(followerId)
	this.createUserIfNotExists(followeeId)

	user := this.users[followerId]
	user.follows[followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}

	this.createUserIfNotExists(followerId)

	user := this.users[followerId]
	delete(user.follows, followeeId)
}

func (this *Twitter) createUserIfNotExists(userId int) {
	if _, ok := this.users[userId]; !ok {
		this.users[userId] = newUser()
	}
}

type maxHeap struct {
	heap [][2]int
}

func newMaxHeap() maxHeap {
	return maxHeap{
		heap: make([][2]int, 0, 10),
	}
}

func (mh *maxHeap) PushTweets(tweets *list.List) {
	for tweet := tweets.Front(); tweet != nil; tweet = tweet.Next() {
		mh.Push(tweet.Value.([2]int))
	}
}

func (mh *maxHeap) ExtractTweets() []int {
	result := make([]int, 0, 10)

	for mh.Len() != 0 {
		tweet := mh.Pop()
		result = append(result, tweet[1])

		if len(result) == 10 {
			break
		}
	}

	return result
}

func (mh *maxHeap) Push(tweet [2]int) {
	mh.heap = append(mh.heap, tweet)
	mh.heapifyUp()
}

func (mh *maxHeap) heapifyUp() {
	index := mh.Len() - 1
	parent := mh.parent(index)

	for mh.heap[parent][0] < mh.heap[index][0] {
		mh.swap(index, parent)
		index = parent
		parent = mh.parent(index)
	}
}

func (mh *maxHeap) Pop() [2]int {
	top := mh.heap[0]

	mh.heap[0] = mh.heap[mh.Len()-1]
	mh.heap = mh.heap[:mh.Len()-1]

	mh.heapifyDown()

	return top
}

func (mh *maxHeap) heapifyDown() {
	index, lastIndex := 0, mh.Len()-1
	left, right := mh.left(index), mh.right(index)

	for left <= lastIndex {
		childToCompare := 0
		if left == lastIndex { // no right
			childToCompare = left
		} else if mh.heap[left][0] > mh.heap[right][0] { // left > right
			childToCompare = left
		} else { // right > left
			childToCompare = right
		}

		if mh.heap[childToCompare][0] > mh.heap[index][0] {
			mh.swap(childToCompare, index)
			index = childToCompare
			left, right = mh.left(index), mh.right(index)
		} else {
			return
		}
	}
}

func (mh *maxHeap) swap(i, j int) {
	mh.heap[i], mh.heap[j] = mh.heap[j], mh.heap[i]
}

func (mh *maxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (mh *maxHeap) left(i int) int {
	return (i * 2) + 1
}

func (mh *maxHeap) right(i int) int {
	return (i * 2) + 2
}

func (mh *maxHeap) Len() int {
	return len(mh.heap)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
