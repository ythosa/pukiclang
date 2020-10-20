package evaluator

import (
    "github.com/ythosa/pukiclang/src/ast"
    "github.com/ythosa/pukiclang/src/object"
)

var (
    NULL  = &object.Null{}
    TRUE  = &object.Boolean{Value: true}
    FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {
    switch node := node.(type) {
    // Statements
    case *ast.Program:
        return evalStatements(node.Statements)

    // Expressions
    case *ast.ExpressionStatement:
        return Eval(node.Expression)

    case *ast.IntegerLiteral:
        return &object.Integer{
            Value: node.Value,
        }

    case *ast.Boolean:
        return nativeBoolToBooleanObject(node.Value)
    }

    return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
    var result object.Object

    for _, statement := range stmts {
        result = Eval(statement)
    }

    return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
    if input {
        return TRUE
    }
    return FALSE
}
