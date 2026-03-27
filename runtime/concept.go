package runtime

/*
	Right now, we're just looking up one dimensional tasks, but
	don't consider if there are nested tasks inside, similar to something
	like this

		IOMeta func do_something(...) (result, error){
			seed := random_gen()
			[[  TimerMeta fn = func () {
				timer := time.NewTimer(seed * time.Millisecond)
				<-timer.C
				run_some_stuff()
			}  ]] >> timerQueue
		}

	But now, you're left with functions that need order of execution
	solely determined by who executes earlier, and you'd now have to track
	them, and their possible ancestors, if they depende on each others return values


	While I think this's interesting, it's out of the scope,like imaginably tons of work
	but for now, think it's best to just do the one-dimensional task metas
*/
type subtask interface {
	parent() fn
}

type task interface {
	Meta
	execute() error
	others() []*Task
}
