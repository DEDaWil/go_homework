package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var textEng = `This tale grew in the telling, until it became a history of the Great War of the Ring and included 
	many glimpses of the yet more ancient history that preceded it. It was begun soon afterThe Hobbit was written 
	and before its publication in 1937; but I did not go on with this sequel, for I wished first to complete and 
	set in order the mythology and legends of the Elder Days, which had then been taking shape for some years. I 
	desired to do this for my own satisfaction, and I had little hope that other people would be interested in 
	this work, especially since it was primarily linguistic in inspiration and was begun in order to provide the 
	necessary background of 'history' for Elvish tongues.

	When those whose advice and opinion I sought correctedlittle hope tono hope, I went back to the sequel, 
	encouraged by requests from readers for more information concerning hobbits and their adventures. But the 
	story was drawn irresistibly towards the older world, and became an account, as it were, of its end and 
	passing away before its beginning and middle had been told. The process had begun in the writing ofThe Hobbit, 
	in which there were already some references to the older matter: Elrond, Gondolin, the High-elves, and the orcs, 
	as well as glimpses that had arisen unbidden of things higher or deeper or darker than its surface: Durin, Moria, 
	Gandalf, the Necromancer, the Ring. The discovery of the significance of these glimpses and of their relation to 
	the ancient histories revealed the Third Age and its culmination in the War of the Ring.`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		expected := []string{
			"а",         // 8
			"он",        // 8
			"и",         // 6
			"ты",        // 5
			"что",       // 5
			"в",         // 4
			"его",       // 4
			"если",      // 4
			"кристофер", // 4
			"не",        // 4
		}
		require.Equal(t, expected, Top10(text))
		expectedEng := []string{
			"the", // 23
			"and", // 14
			"of",  // 11
			"in",  // 9
			"i",   // 6
			"to",  // 6
			"for", // 5
			"had", // 5
			"it",  // 5
			"its", // 5
		}
		require.Equal(t, expectedEng, Top10(textEng))
	})
}
