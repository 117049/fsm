package main

import (
	"fmt"
	"strings"
	"time"
)

/*
 * author : xhs
 * date : 2022/6/12
 * describe : 状态机测试实现,
	状态： 吃饭，睡觉，游戏
	事件驱动： 饿了， 困了， 无聊
	三种时间驱动三种状态的转换，但是将状态切换的对象交给了状态本身，不易于动态扩展，可是通过状态机进行状态的转换
*/

type STATE_TYPE int

// 状态
const (
	STATE_SLEEPING STATE_TYPE = iota
	STATE_EATING
	STATE_GAMING
	STATE_END
)

type EVENT_TYPE int

// 事件
const (
	EVENT_SLEEPY EVENT_TYPE = iota
	EVENT_HUNGARY
	EVENT_BORING
	EVENT_END
)

// 状态类，包含进入，离开，维持
type STATE interface {
	enter()
	exit()
	stay()
	eventHandle(*STATE_MACHION, EVENT_TYPE)
}

// sleeping状态
type statesleeping struct {
	sleepState bool
}

func (state *statesleeping) enter() {
	fmt.Println("sleeping enter")
	state.sleepState = true
	state.stay()
}

func (state *statesleeping) exit() {
	fmt.Println("sleeing exit")
	state.sleepState = false
}

func (state *statesleeping) stay() {
	go func() {
		for {
			if state.sleepState == false {
				break
			}
			fmt.Println("sleeing stay")
			time.Sleep(5 * time.Second)
		}
	}()
}

func (state *statesleeping) eventHandle(sm *STATE_MACHION, event EVENT_TYPE) {
	switch event {
	case EVENT_HUNGARY:
		sm.switchToState(STATE_EATING)
		break
	case EVENT_BORING:
		sm.switchToState(STATE_GAMING)
	}

}

// eating状态
type stateeating struct {
	eatingstate bool
}

func (state *stateeating) enter() {
	fmt.Println("stateeating enter")
	state.eatingstate = true
}
func (state *stateeating) exit() {
	fmt.Println()
	state.eatingstate = false
}
func (state *stateeating) stay() {
	go func() {
		for {
			if state.eatingstate == false {
				break
			}
			fmt.Println("eatingstate stay")
			time.Sleep(5 * time.Second)
		}
	}()

}

func (state *stateeating) eventHandle(sm *STATE_MACHION, event EVENT_TYPE) {
	switch event {
	case EVENT_SLEEPY:
		sm.switchToState(STATE_SLEEPING)
		break
	case EVENT_BORING:
		sm.switchToState(STATE_GAMING)
	}
}

// gaming状态
type stategaming struct {
	gamingstate bool
}

func (state *stategaming) enter() {
	fmt.Println("gaming enter")
	state.gamingstate = true
}
func (state *stategaming) exit() {
	fmt.Println("gaming exit")
	state.gamingstate = false
}
func (state *stategaming) stay() {
	go func() {
		for {
			if state.gamingstate == false {
				break
			}
			fmt.Println("gaming stay")
			time.Sleep(5 * time.Second)
		}
	}()
}
func (state *stategaming) eventHandle(sm *STATE_MACHION, event EVENT_TYPE) {
	switch event {
	case EVENT_SLEEPY:
		sm.switchToState(STATE_SLEEPING)
		break
	case EVENT_HUNGARY:
		sm.switchToState(STATE_EATING)
	}
}

// 状态机类
type STATE_MACHION struct {
	curStateType STATE_TYPE
	statePool    map[STATE_TYPE]STATE
}

func (sm *STATE_MACHION) addState(state_type STATE_TYPE, state STATE) {
	sm.statePool[state_type] = state
}

func (sm *STATE_MACHION) initState(state_type STATE_TYPE) {
	sm.curStateType = state_type
	sm.statePool[state_type].enter()
}
func (sm *STATE_MACHION) switchToState(state_type STATE_TYPE) {
	sm.statePool[sm.curStateType].exit()
	sm.statePool[state_type].enter()
	sm.curStateType = state_type
}
func (sm *STATE_MACHION) eventHandle(event EVENT_TYPE) {
	sm.statePool[sm.curStateType].eventHandle(sm, event)
}
func main() {
	// 构造对象
	staSleep := &statesleeping{}
	staEat := &stateeating{}
	staGam := &stategaming{}

	sm := &STATE_MACHION{}

	// 添加状态
	sm.statePool = make(map[STATE_TYPE]STATE)
	sm.addState(STATE_SLEEPING, staSleep)
	sm.addState(STATE_EATING, staEat)
	sm.addState(STATE_GAMING, staGam)

	sm.initState(STATE_SLEEPING)

	var inputChar string
	for {
		inputChar = ""
		fmt.Scanln(&inputChar)
		fmt.Println(inputChar)
		if strings.Compare(inputChar, "sleep") == 0 {
			sm.eventHandle(EVENT_SLEEPY)
		} else if strings.Compare(inputChar, "hungry") == 0 {
			sm.eventHandle(EVENT_HUNGARY)
		} else if strings.Compare(inputChar, "boring") == 0 {
			sm.eventHandle(EVENT_BORING)
		}
	}
}
