package entities

type Scene int

const (
	SceneAddCalculation Scene = iota + 1
	SceneAddBill
	SceneAddDebt
	SceneAddComplexBill
)

type Question int

const (
	QuestionBillDesc Question = iota + 1
	QuestionBillPrice
	QuestionBillPayer
	QuestionBillDebtor
)

const (
	QuestionDebtDesc Question = iota + 1
	QuestionDebtPrice
	QuestionDebtPayer
	QuestionDebtDebtor
)

const (
	QuestionComplexBillDesc Question = iota + 1
	QuestionComplexBillPrice
	QuestionComplexBillPayer
	QuestionComplexBillForm

	QuestionComplexBillEntry
	QuestionComplexBillService
)

func (q Question) NextQuestion() Question {
	return q + 1
}
