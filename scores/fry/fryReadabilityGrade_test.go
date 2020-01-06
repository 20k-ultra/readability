package fry

import (
	"testing"
)

var PrivacyPolicy = []byte(`
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
		{name: " Calculate Fry Readability standard", args: args{text: ShortStory}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.args.text); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
