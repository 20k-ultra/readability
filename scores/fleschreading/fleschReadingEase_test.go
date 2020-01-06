package fleschreading

import "testing"

var ShortStory = []byte(`Day had broken cold and grey, exceedingly cold and grey, when the man turned aside from the main Yukon trail and climbed the high earth- bank, where a dim and little-travelled trail led eastward through the fat spruce timberland. It was a steep bank, and he paused for breath at the top, excusing the act to himself by looking at his watch. It was nine o'clock. There was no sun nor hint of sun, though there was not a cloud in the sky. It was a clear day, and yet there seemed an intangible pall over the face of things, a subtle gloom that made the day dark, and that was due to the absence of sun. This fact did not worry the man. He was used to the lack of sun. It had been days since he had seen the sun, and he knew that a few more days must pass before that cheerful orb, due south, would just peep above the sky- line and dip immediately from view.`)
var LegalText = []byte(`BINDING EFFECT This Agreement shall be binding upon the parties thereto and their respective legal representatives, successors and assignees. SEVERABILITY
If any provision of this agreement is invalid or unenforceable, the balance of this agreement shall remain in effect, and if any provision is inapplicable to any person or circumstances, it shall nevertheless remain applicable to all other persons and circumstances.
IN WITNESS WHEREOF, the parties hereto have read and fully understood its contents and hereby agree to comply with its terms and conditions and have caused this Agreement to be duly executed by their authorized representatives affixed with the seals thereon in the presence of witnesses on the date first above written.`)
var PrivacyPolicy = []byte(`We only collect the information you choose to give us, and we process it with your consent, or on another legal basis; we only require the minimum amount of personal information that is necessary to fulfill the purpose of your interaction with us; we don't sell it to third parties; and we only use it as this Privacy Statement describes. If you're visiting us from the EU, please see our global privacy practices: we comply with the Privacy Shield framework and we are compliant with the General Data Protection Regulation (GDPR). No matter where you are, where you live, or what your citizenship is, we provide the same standard of privacy protection to all our users around the world, regardless of their country of origin or location.`)
var DaftPunkLyrics = []byte(`Around the world, around the world`)

func TestCalculate(t *testing.T) {

	type args struct {
		text []byte
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{name: " Calculate fleschreading standard", args: args{text: DaftPunkLyrics}, want: 87.95},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.args.text); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
