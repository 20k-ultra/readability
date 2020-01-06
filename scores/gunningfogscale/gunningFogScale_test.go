package gunningfogscale

import "testing"

var ShortStory = []byte(`As a single working mom her life was a bit harried but she took it all in stride. Life was good and she loved her son John more than anything else in the world. He was her raison d’être (ray-zohnn DET-ruh).
When John was just an infant he sucked his thumb. By age two he became attached to the silky edge of his baby blanket. Rather than carry that heavy blanket everywhere mom made silky toy by wrapping silky white fabric around a soft rubber ball and securing it with a rubber band (like a Halloween decoration.)
John immediately became attached to his “Ghostie.”
One advantage to Ghostie was he could easily be remade in case he disappeared.
When John was about three years old mom made plans to eat at her favorite restaurant Michael J’s. It was not often they could afford to go out but you have to treat yourself once in awhile. She made sure John was well rested so there would be no fussiness in public. John invited Ghostie to dinner and Ghostie accepted of course. Even though he was never hungry.
It was a fine meal – cheeseburgers and ice cream all around.
With bellies full and the sun sinking below the horizon, mom strapped John into his car seat and home they went.
About five miles down the road John asked, “where’s Ghostie?”
“Oh no! We must have left him at Michael J’s!”
Mom quickly turned the car around.
Back at the restaurant they both marched up to the man standing at the little podium and asked in unison, “have you seen a Ghost?”
As the blood drained from the man’s face a waitress overheard and came running over Ghostie in hand. “I knew you’d be back,” she said with a smile.
 The End.`)

func TestCalculate(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{name: " Calculate Gunning Fog Scale standard", args: args{text: ShortStory}, want: 7.7 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.args.text); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
