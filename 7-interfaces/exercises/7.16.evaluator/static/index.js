function Expression(init) {
	let values = init || { args: [] }

	let evaluated = false
	let pending = false

	return {
		storeOperator,
		storeOperand,
		storeFunction,
		updateOperand,
		evaluated,
		calculate,
		operator,
		operands,
		pending,
		reset
	}

	function storeOperand(name) {
		values.args.push(name)
	}

	function storeOperator(name) {
		values.op = name
	}

	function storeFunction(name) {
		values.fun = name
	}

	function updateOperand(index, value) {
		values.args[index] = value
	}

	async function calculate() {
		let expression
		switch (values.op) {
			case '+':
				expression = `${values.args[0]}%2b${values.args[1]}`
				break
			case '-':
				expression = `${values.args[0]}-${values.args[1]}`
				break
			case '*':
				expression = `${values.args[0]}*${values.args[1]}`
				break
			case '/':
				expression = `${values.args[0]}/${values.args[1]}`
				break
			case 'pow':
				expression = `pow(${values.args[0]},${values.args[1]})`
				break
		}
		switch (values.fun) {
			case 'sqrt':
				expression = `sqrt(${values.args[0]})`
				break
			case 'sin':
				expression = `sin(${values.args[0]})`
				break
		}

		// handle single unary operator expressions, e.g., -x
		if (!expression) {
			expression = values.args.reverse().join()
		}

		response = await fetch("calculate?expr=" + expression)
		return await response.text()
	}

	function operands() {
		return values.args
	}

	function operator() {
		return values.op
	}

	function reset() {
		values = { args: [] }
	}
}

function Calculator(expression) {
	var value = "0"

	setReadout(value)

	return {
		set: setReadout,
		save: saveReadout,
		clear: clearReadout,
		execute: runFunction,
		calculate: runOperator,
		operator: expression.operator,
		operands: expression.operands
	}


	function setReadout(operand) {
		document.getElementsByClassName("clear")[0].innerText = "C"
		document.getElementsByClassName("clear")[0].onclick = clearReadout

		let readout = document.getElementsByTagName("input")[0]

		// ignore adding unary to zero value expressions
		if (readout.value == "0" && operand == "-") {
			return
		}

		// avoid leading zero when setting numbers
		if (readout.value == "0") {
			readout.value = operand
			return
		}

		// toggle unary operand
		if (operand == "-" && parseFloat(readout.value) < 0) {
			readout.value = JSON.stringify(Math.abs(readout.value))
			return
		}

		// make sure unary is ordered before operator
		if (operand == '-') {
			readout.value = operand + readout.value
			return
		}

		// when an expression has been evaluated (e.g. i + n = x) or the
		// expression is in a pending state (e.g. i + n where n is the last
		// selected operator). This allows us to clear the readout when either
		// states are true.
		if (expression.evaluated || expression.pending) {
			readout.value = operand
			expression.evaluated = false
			if (expression.pending) {
				readout.value = operand
				expression.pending = false
				return
			}
			return
		}

		readout.value = readout.value + operand
	}

	async function saveReadout(operator) {
		let readout = document.getElementsByTagName("input")[0]

		// evaluate existing expression before continuing, e.g., (n + n) + n
		if (expression.operator()) {
			expression.updateOperand(1, readout.value)
			readout.value = await expression.calculate()
			expression.reset()
		}

		let operand = readout.value
		expression.storeOperand(operand)
		expression.storeOperator(operator)
		expression.pending = true
	}

	function clearReadout() {
		document.getElementsByTagName("input")[0].value = 0
		document.getElementsByClassName("clear")[0].innerText = "AC"
		document.getElementsByClassName("clear")[0].onclick = clearAllReadout
		expression.evaluated = false
	}

	function clearAllReadout() {
		document.getElementsByTagName("input")[0].value = 0
		expression.pending = false
		expression.reset()
	}

	async function runFunction(fun) {
		let readout = document.getElementsByTagName("input")[0]

		expression.storeOperand(readout.value)
		expression.storeFunction(fun)

		readout.value = await expression.calculate()
		expression.reset()
	}

	async function runOperator() {
		let readout = document.getElementsByTagName("input")[0]

		// support unary and literal expressions without operators
		if (readout.value < 0 || expression.operands().length == 0) {
			expression.storeOperand(readout.value)
			readout.value = await expression.calculate()
			expression.reset()
			return
		}

		expression.updateOperand(1, readout.value) // account for multi-digit
		readout.value = await expression.calculate()
		expression.evaluated = true
		expression.reset()
	}
}

var calculator = new Calculator(new Expression)
