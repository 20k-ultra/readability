package readability

import (
	"reflect"
	"testing"
)

var ShortStory = []byte(`Day had broken cold and grey, exceedingly cold and grey, when the man turned aside from the main Yukon trail and climbed the high earth- bank, where a dim and little-travelled trail led eastward through the fat spruce timberland. It was a steep bank, and he paused for breath at the top, excusing the act to himself by looking at his watch. It was nine o'clock. There was no sun nor hint of sun, though there was not a cloud in the sky. It was a clear day, and yet there seemed an intangible pall over the face of things, a subtle gloom that made the day dark, and that was due to the absence of sun. This fact did not worry the man. He was used to the lack of sun. It had been days since he had seen the sun, and he knew that a few more days must pass before that cheerful orb, due south, would just peep above the sky- line and dip immediately from view.`)
var LegalText = []byte(`BINDING EFFECT This Agreement shall be binding upon the parties thereto and their respective legal representatives, successors and assignees. SEVERABILITY
If any provision of this agreement is invalid or unenforceable, the balance of this agreement shall remain in effect, and if any provision is inapplicable to any person or circumstances, it shall nevertheless remain applicable to all other persons and circumstances.
IN WITNESS WHEREOF, the parties hereto have read and fully understood its contents and hereby agree to comply with its terms and conditions and have caused this Agreement to be duly executed by their authorized representatives affixed with the seals thereon in the presence of witnesses on the date first above written.`)
var PrivacyPolicy = []byte(`We only collect the information you choose to give us, and we process it with your consent, or on another legal basis; we only require the minimum amount of personal information that is necessary to fulfill the purpose of your interaction with us; we don't sell it to third parties; and we only use it as this Privacy Statement describes. If you're visiting us from the EU, please see our global privacy practices: we comply with the Privacy Shield framework and we are compliant with the General Data Protection Regulation (GDPR). No matter where you are, where you live, or what your citizenship is, we provide the same standard of privacy protection to all our users around the world, regardless of their country of origin or location.`)
var pieceOfPrivacyPolicy= []byte(`No matter where you are, where you live, or what your citizenship is, we provide the same standard of privacy protection to all our users around the world, regardless of their country of origin or location.`)
var DaftPunkLyrics = []byte(`Around the world, around the world`)
var smallByteSlice = []byte(`hey my guy whatsup.`)
var LongPrivacyPolicy = []byte(`
Like most websites, our servers automatically record the page requests made when you visit our sites. These “server logs” typically include your web request, Internet Protocol address, browser type, browser language, the date and time of your request, and one or more cookies that may uniquely identify your browser.

123.45.67.89 is the Internet Protocol address assigned to the user by the user’s ISP. Depending on the user’s service, a different address may be assigned to the user by their service provider each time they connect to the Internet.
25/Mar/2003 10:15:32 is the date and time of the query.
http://www.google.com/search?q=cars is the requested URL, including the search query.
Firefox 1.0.7; Windows NT 5.1 is the browser and operating system being used.
740674ce2123a969 is the unique cookie ID assigned to this particular computer the first time it visited Google. (Cookies can be deleted by users. If the user has deleted the cookie from the computer since the last time they’ve visited Google, then it will be the unique cookie ID assigned to their device the next time they visit Google from that particular device).

Unique identifiers

A unique identifier is a string of characters that can be used to uniquely identify a browser, app, or device. Different identifiers vary in how permanent they are, whether they can be reset by users, and how they can be accessed.

Unique identifiers can be used for various purposes, including security and fraud detection, syncing services such as your email inbox, remembering your preferences, and providing personalized advertising. For example, unique identifiers stored in cookies help sites display content in your browser in your preferred language. You can configure your browser to refuse all cookies or to indicate when a cookie is being sent. Learn more about how Google uses cookies.

On other platforms besides browsers, unique identifiers are used to recognize a specific device or app on that device. For example, a unique identifier such as the Advertising ID is used to provide relevant advertising on Android devices, and can be managed in your device’s settings. Unique identifiers may also be incorporated into a device by its manufacturer (sometimes called a universally unique ID or UUID), such as the IMEI-number of a mobile phone. For example, a device’s unique identifier can be used to customize our service to your device or analyze device issues related to our services.
`)

func TestGetSentences(t *testing.T) {
	type args struct {
		textBytes []byte
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Get Sentences standard", args: args{textBytes: DaftPunkLyrics}, want: []string{
			"Around the world, around the world",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSentences(tt.args.textBytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWords(t *testing.T) {
	type args struct {
		text          []byte
		removePeriods bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Get Words standard", args: args{text: DaftPunkLyrics, removePeriods: false}, want: []string{
			"Around", "the", "world", "around", "the", "world",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWords(tt.args.text, tt.args.removePeriods); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountCharactersPerWord(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{name: "Get Characters Per Word standard", args: args{words: GetWords(DaftPunkLyrics, false)}, want: 4.67},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountCharactersPerWord(tt.args.words); got != tt.want {
				t.Errorf("countCharactersPerWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountWordsPerSentence(t *testing.T) {
	type args struct {
		sentences []string
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{name: "Get Words Per Sentence standard", args: args{sentences: GetSentences(DaftPunkLyrics)}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountWordsPerSentence(tt.args.sentences); got != tt.want {
				t.Errorf("countWordsPerSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountSyllablesPerWord(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{name: "Get Syllables Per Word standard", args: args{words: GetWords(LegalText, false)}, want: 1.8849558},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSyllablesPerWord(tt.args.words); got != tt.want {
				t.Errorf("CountSyllablesPerWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountSyllables(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Get Syllables In Word standard", args: args{word: "Around"}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSyllables(tt.args.word); got != tt.want {
				t.Errorf("CountSyllablesInWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountSyllablesInSlice(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Get Number of Syllables In Words standard", args: args{words: GetWords(DaftPunkLyrics, false)}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSyllablesInSlice(tt.args.words); got != tt.want {
				t.Errorf("GetNumberOfSyllables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWordSample(t *testing.T) {
	type args struct {
		words      []string
		sampleSize int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: " Get Word Sample standard",
			args: args{words: GetWords(smallByteSlice, false), sampleSize: 2},
			want: []string{
				"hey my", "guy whatsup.",
			},
		},
		// {	name: " Get Word Sample standard #2",
		// 	args: args{words: GetWords(LongPrivacyPolicy, true), sampleSize: 2},
		// 	want: []string{
		// 		"hey my", "guy.",
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWordSample(tt.args.words, tt.args.sampleSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getWordSamples() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcOffsetIndexes(t *testing.T) {
	type args struct {
		size   int
		length int
	}
	tests := []struct {
		name string
		args args
		want []offset
	}{
		{name: " Calculate Offset Indexes standard",
			args: args{size: len(GetWords(smallByteSlice, false)), length: 2},
			want: []offset{
				offset{start: 0, end: 1},
				offset{start: 2, end: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcOffsetIndexes(tt.args.size, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcOffsetIndexes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWordsInRange(t *testing.T) {
	type args struct {
		words []string
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Get Words In Range standard",
			args: args{words: GetWords(smallByteSlice, true), start: 0, end: 1},
			want: "hey my"},
		{name: "Get Words In Range standard 2",
			args: args{words: GetWords(smallByteSlice, false), start: 2, end: 3},
			want: "guy whatsup."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWordsInRange(tt.args.words, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("GetWordsInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
