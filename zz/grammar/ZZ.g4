grammar ZZ;

Var : 'var' ;
Int : 'int' ;
Float : 'float' ;
Matrix : 'matrix' ;

Func : 'func' ;
List : 'list' ;

Nil : 'nil';

Identifier
    :   [a-zA-Z] [a-zA-Z_0-9]*
    ;

IntegerLiteral
    :   '0'
    |   [1-9] [0-9]*
    ;

FloatLiteral // todo
    :   '0' '.' [0-9]*
    |   [1-9] [0-9]* '.'? [0-9]*
    |   '-' FloatLiteral
    ;

identifier
    :   Identifier
    ;

declarator
    :   identifier
    |   listElementExpression
    ;

declaratorList
    :   declarator
    |   declaratorList ',' declarator
    ;

simpleTypeSpecifier
    :   'int'
    |   'float'
    |   'string'
    |   'matrix'
    ;

listElementTypeSpecifier
    :   simpleTypeSpecifier
    |   '[' ']' listElementTypeSpecifier
    ;

listTypeSpecifier
    :   '[' ']' listElementTypeSpecifier
    ;

listInitExpression
    :   'list' '(' listTypeSpecifier ',' aExpr ')'
    ;

matrixInitExpression
    :   'matrix' '(' aExprList ')'
    ;

aExpr
    :   IntegerLiteral # aExp_IntergerLiteral
    |   FloatLiteral # aExp_FloatLiteral
    |   identifier # aExp_identifier
    |   listElementExpression # aExp_listElementExpression
    |   aExpr ('*'|'/') aExpr # aExp_multiplicativeExpression
    |   aExpr ('+'|'-') aExpr # aExp_additiveExpression
    |   '(' aExpr ')' # aExp_bracketExpression
    |   'transpose' '(' aExpr ')' # aExp_transpose
    ;

aExprList
    :   aExpr
    |   aExprList ',' aExpr
    ;

bExpr
    :   aExpr ('=='|'<'|'>'|'<='|'>='|'!=') aExpr # bExpr_aExpr
    |   bExpr ('=='|'&&'|'||'|'!=') bExpr # bExpr_bExpr
    |   '!' bExpr # bExpr_bang
    |   '(' bExpr ')' # bExpr_bracketExpression
    ;

listElementIndex
    :   '[' aExpr ']'
    ;

listElementIndexList
    :   listElementIndex
    |   listElementIndexList listElementIndex
    ;

listElementExpression
    :   identifier listElementIndexList
    ;

assignInit
    :   aExpr
    |   listInitExpression
    |   matrixInitExpression
//    |   funcInitExpression
    |   funcExecuteExpression
    ;

assignInitList
    :   assignInit
    |   assignInitList ',' assignInit
    ;

assignStatement
    :   declaratorList '=' assignInitList
    ;

assignInitStatement
    :   declaratorList ':=' assignInitList
    ;

ifExpr
    :   'if' bExpr '{' funcStatementList? '}'
    ;

elsifExpr
    :   'elsif' bExpr '{' funcStatementList? '}'
    ;

elseExpr
    :   'else' '{' funcStatementList? '}'
    ;

ternaryIfExpr
    :   bExpr '?' funcStatement
    ;

ternaryElseExpr
    :   funcStatement
    ;

selectionStatement
    :   ifExpr elsifExpr* elseExpr?
    |   ternaryIfExpr ':' ternaryElseExpr
    ;

iterationAssignInitStatement
    :   assignInitStatement
    |   assignStatement
    ;

iterationAssignStatement
    :   assignStatement
    ;

iterationStatement
    :   'for' iterationAssignInitStatement? ';' bExpr? ';' iterationAssignStatement? '{' funcStatementList? '}' // todo
    ;

definition
    :   assignInitStatement
    |   funcDefinition
    ;

definitionList
    :   definition
    |   definitionList definition
    ;

file // todo: package import
    :   definitionList
    ;

typeSpecifier
    :   simpleTypeSpecifier
    |   listTypeSpecifier
//    |   funcTypeSpecifier
    ;

typeSpecifierList
    :   typeSpecifier
    |   typeSpecifierList ',' typeSpecifier
    ;

paraDeclarator
    :   identifier
    ;

paraDeclaratorList
    :   paraDeclarator
    |   paraDeclaratorList ',' paraDeclarator
    ;

paraDeclaratorWithIdentity // todo: rename
    :   paraDeclaratorList typeSpecifier
    ;

paraDeclaratorWithIdentityList
    :   paraDeclaratorWithIdentity
    |   paraDeclaratorWithIdentityList ',' paraDeclaratorWithIdentity
    ;

//funcTypeSpecifier
//    :   'func' '(' paraDeclaratorWithIdentityList? ')' ('(' typeSpecifierList? ')')?
//    ;

funcIdentifier
    :   identifier
    ;

funcTypeSpecifierWithName
    :   'func' funcIdentifier '(' paraDeclaratorWithIdentityList? ')' ('(' typeSpecifierList? ')')?
    ;

funcReturnPara
    :   aExpr
    |   bExpr
    |   identifier
    |   'nil'
    |   funcExecuteExpression
    ;

funcReturnParaList
    :   funcReturnPara
    |   funcReturnParaList ',' funcReturnPara
    ;

funcReturnStatement
    :   'return' funcReturnParaList?
    ;

funcStatement
    :   assignStatement
    |   assignInitStatement
    |   selectionStatement
    |   iterationStatement
    |   funcReturnStatement
    |   funcExecuteStatement
    |   printStatement
    ;

funcStatementList
    :   funcStatement
    |   funcStatementList funcStatement
    ;

funcBody
    :   funcStatementList?
    ;

//funcInitExpression // when assign to a varient: function = func() {}
//    :   funcTypeSpecifier '{' funcBody? '}'
//    ;

funcDefinition // when define a method for a class or a static function: func function() {}
    :   funcTypeSpecifierWithName '{' funcBody? '}'
    ;

funcExecutePara
    :   aExpr
    |   bExpr
    ;

funcExecuteParaList
    :   funcExecutePara
    |   funcExecuteParaList ',' funcExecutePara
    ;

funcExecuteExpression
    :   funcIdentifier '(' funcExecuteParaList? ')'
    ;

funcExecuteStatement
    :   funcExecuteExpression
    ;

fragment
StringChar
    :   [a-zA-Z_0-9 \r\t\n]
    ;

printList
    :   (aExpr | bExpr)
    |   printList ',' ( aExpr | bExpr)
    ;

printStatement
    :   'print' '(' printList ')'
    ;

WS : [ \t\r\n]+ -> skip ; // skip spaces, tabs, newlines
