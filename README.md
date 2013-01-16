# Stack

A basic and very simple implementation of a stack in Go. For more information on what a stack is, read this [wikipedia article](https://en.wikipedia.org/wiki/Stack_\(abstract_data_type\)#Software_stacks).

## Usage

Get the package:

	$ go get github.com/sauerbraten/stack

Import the package:

	import (
		"github.com/sauerbraten/stack"
	)

Using a stack is pretty straightforward:

	s := stack.New()

	// push '3' (int) onto the stack
	s.Push(3)

	// top element is now '3', size is 1
	fmt.Println(s.Size())
	// prints: 1

	// push "four" (string) onto the stack
	s.Push("four")

	// top element is now "four", size is 2
	fmt.Println(s.Size())
	// prints: 2

	// pop the "four" from the stack
	fmt.Println(s.Pop())
	// prints: four <nil>

	// now, the top element is '3'

	// pop the '3' from the stack
	fmt.Println(s.Pop())
	// prints: 3 <nil>

	// stack is now empty
	fmt.Println(s.Pop())
	// prints: <nil> stack is empty

	fmt.Println(s.Size())
	// prints: 0

## Documentation

For full package documentation, visit http://godoc.org/github.com/sauerbraten/stack.

## License

This code is licensed under a BSD License:

	Copyright (c) 2012 Alexander Willing. All rights reserved.
		* Redistributions of source code must retain the above copyright
		  notice, this list of conditions and the following disclaimer.
		* Redistributions in binary form must reproduce the above
		  copyright notice, this list of conditions and the following
		  disclaimer in the documentation and/or other materials provided
		  with the distribution.

	THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
	"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
	LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
	A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
	OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
	SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
	LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
	DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
	THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
	(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
	OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
