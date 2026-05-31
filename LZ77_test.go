package lz77

import (
	"testing"
	"os"
)

const lorem string = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec quis libero ac arcu sodales laoreet quis vitae enim. Sed placerat arcu nec lorem consequat pellentesque. Cras accumsan velit sed fermentum suscipit. Pellentesque efficitur libero eu faucibus pulvinar. Curabitur aliquam libero fermentum tempor posuere. Cras euismod odio ac sagittis ornare. Aliquam vel venenatis neque. Donec at posuere justo. Donec iaculis orci eu hendrerit dignissim. Nulla facilisi. Praesent ac purus quis nisi facilisis posuere. Curabitur at urna vitae justo laoreet ornare nec vel tellus. Curabitur accumsan justo pulvinar, aliquet odio ut, auctor metus. Mauris lectus libero, malesuada sed sem non, egestas rhoncus urna. Aenean ipsum mauris, commodo in suscipit sit amet, tincidunt non sapien. Aenean a leo tortor.Nullam quis erat quam. Etiam a sem ut odio congue suscipit ut eget neque. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam eu vehicula purus, ut semper ligula. Mauris bibendum lacinia nisi sed volutpat. Quisque vel dolor elit. Nulla facilisi. Nullam nec ex ut odio elementum laoreet. Nulla ut ultrices velit, a gravida tortor. Mauris eget tristique tellus. Praesent mollis, odio non facilisis cursus, tellus ex eleifend nunc, et consectetur lectus nisl ac ipsum.Morbi mi nunc, viverra eu pretium non, rutrum ac magna. Donec interdum pulvinar est, eget luctus lacus commodo eu. Morbi tempor dignissim tellus, vel interdum lectus auctor eu. Interdum et malesuada fames ac ante ipsum primis in faucibus. Nam accumsan vestibulum odio sed consequat. Mauris eu placerat nulla. Integer consectetur in justo quis lacinia. Fusce maximus purus enim, eget tempor odio vulputate sed. Aenean commodo massa massa, mattis vestibulum dui dapibus in. Pellentesque porta fringilla nunc, rhoncus imperdiet turpis lobortis in.Morbi lectus sem, tincidunt ut volutpat ac, congue eu nunc. Morbi feugiat odio sem, vel egestas orci dictum eget. Sed hendrerit massa sit amet posuere pharetra. Nunc eleifend gravida dui, eu congue nisi porttitor id. Pellentesque eu leo id massa placerat rhoncus. Vivamus euismod dui eu leo venenatis, ac viverra metus convallis. Curabitur convallis sollicitudin ante vel condimentum. Morbi non maximus ante. Sed dictum imperdiet gravida. Duis lectus justo, tristique a tempus laoreet, egestas ut nisi. In nisi enim, ornare sit amet ligula in, ullamcorper tempor purus.Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Nulla egestas tellus quis leo fringilla lacinia. Nullam ut leo fringilla, fermentum massa vel, tempor enim. Etiam pretium mi at accumsan varius. Duis congue efficitur justo fringilla varius. In id erat suscipit, pellentesque nisl vel, vulputate nulla. Suspendisse potenti. Sed vel aliquam diam. Etiam mattis commodo odio, vel eleifend dui commodo sit amet. Quisque ac condimentum arcu. Donec sollicitudin vitae tortor id egestas. Mauris consectetur libero vitae euismod pulvinar. Quisque tempus porta interdum. Cras mollis congue sem, ac rhoncus est facilisis at. Praesent maximus est eu neque commodo, id egestas magna aliquet."

func TestLz(t *testing.T) {
	c := Compress(lorem, 32000)
	res := Decompress(c)

	if res != lorem {
		t.Fatal("Compression failed")
	}
}

func TestWriteFile(t *testing.T) {
	test := "tres tristes tigres tragaban trigo en un trigal"
	c := Compress(test, 32000)
	err := WriteResultFile("test.lz77", c)

	if err != nil {
		t.Fatal("WriteFile failed")
	}
	os.Remove("test.lz77")
}

func TestReadFile(t *testing.T) {
	test_content := "yipiyakei"
	err := os.WriteFile("test.input", []byte(test_content), 0644)
	if err != nil {
		t.Fatal("Could not write input test file", err)
	}

	read, error := ReadFile("test.input")
	if error != nil || string(read) != test_content {
		t.Fatal("Readed content is not equal to test content")
	}
	os.Remove("test.input")
}

func TestResultFromBytes(t *testing.T) {
	c := Compress(lorem, 32000)
	
	err := WriteResultFile("test.lz77", c)
	if err != nil {
		t.Fatal("ResultFromBytes WriteFile failed")
	}

	read, err := ReadFile("test.lz77")
	if err != nil {
		t.Fatal("ResultFromBytes ReadFile error")
	}

	parsed := ParseBytes(read)
	if len(parsed) != len(c) {
		t.Fatal("Parsed not equal than compressed")
	}
	os.Remove("test.lz77")
}
