package enums

// GameMode is a flag for the game mode
type GameMode int8

const (
	GM_EventMode                GameMode = 0   // event mode
	GM_GrandPrix                GameMode = 3   // grand prix
	GM_GrandPrix23              GameMode = 4   // grand prix23
	GM_TimeTrial                GameMode = 5   // time trial
	GM_Splitscreen              GameMode = 6   // split screen
	GM_OnlineCustom             GameMode = 7   // online custom
	GM_OnlineLeague             GameMode = 8   // online league
	GM_CareerInvitational       GameMode = 11  // career invitational
	GM_ChampionshipInvitational GameMode = 12  // championship invitational
	GM_Championship             GameMode = 13  // championship
	GM_OnlineChampionship       GameMode = 14  // online championship
	GM_OnlineWeeklyEvent        GameMode = 15  // online weekly event
	GM_StoryMode                GameMode = 17  // story mode
	GM_Career22                 GameMode = 19  // career 22
	GM_Career22Online           GameMode = 20  // career 22 online
	GM_Career23                 GameMode = 21  // career 23
	GM_Career23Online           GameMode = 22  // career 23 online
	GM_Benchmark                GameMode = 127 // benchmark
)
