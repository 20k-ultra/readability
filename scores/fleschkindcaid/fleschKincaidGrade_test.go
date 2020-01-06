package fleschkindcaid

import "testing"

var story = []byte(`I opened my eyes to see bright blinding light coming through the thin curtains and flooding the entire room. A frustrated groan escaped my lips as I rolled over on my small bed, turning away from the large window and coming face to face with my alarm clock. I jumped out of bed, stumbling on the pallet board that stuck out from beneath my mattress and ran straight to the bathroom down the hall. Work started at nine and I only had a few minutes to get ready. I stood in front of the mirror analysing my appearance. My bright blue eyes were too large on my heart-shaped face and dark semicircles sat beneath my lower eyelids as proof of my late night. My thick dark hair cascaded past my shoulders and down to my waist in a tangled mess. I looked away from my reflection and brushed my teeth in a hurry before changing from my flannel pyjamas, covering my head with a large warm beanie, grabbing a few things and darting out of the house. I made sure to lock all the doors before leaping off the front porch. The icy gravel on the small footpath crunched rhythmically beneath my worn boots as I made my way to my car, a 1948 Ford F1. The rusty pastel coloured truck belonged to my father, it had him written all over it, from the rusty undercarriage to the broken side mirror and the crooked rear tyre. It was kooky and unconventional just like he was.`)

func TestCalculate(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{name: " Calculate fleschkindcaid standard", args: args{text: story}, want: 8.29},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.args.text); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
