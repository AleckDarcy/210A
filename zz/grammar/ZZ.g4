grammar ZZ;

Var : 'var' ;
Int : 'int' ;
Float : 'float' ;

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
    :   '0.0'
    ;

declarator
    :   Identifier
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
    ;

listElementTypeSpecifier
    :   simpleTypeSpecifier
    |   '[' ']' listElementTypeSpecifier
    ;

listTypeSpecifier
    :   '[' ']' listElementTypeSpecifier
    ;

listInitExpression
    :   'list' '(' listTypeSpecifier ',' '(' aExpr ')' ')'
    ;

aExpr
    :   IntegerLiteral # aExp_IntergerLiteral
    |   FloatLiteral # aExp_FloatLiteral
    |   Identifier # aExp_Identifier
    |   listElementExpression # aExp_listElementExpression
    |   aExpr ('*'|'/') aExpr # aExp_multiplicativeExpression
    |   aExpr ('+'|'-') aExpr # aExp_additiveExpression
    |   '(' aExpr ')' # aExp_bracketExpression
    ;

aExprList
    :   aExpr
    |   aExprList ',' aExpr
    ;

bExp
    :   aExpr ('=='|'<'|'>'|'>='|'<='|'!=') aExpr
    |   bExp ('=='|'&&'|'||'|'!=') bExp
    |   '!' bExp
    |   '(' bExp ')'
    ;

integerExpression
    :   IntegerLiteral
    |   Identifier
    ;

listElementIndex
    :   '[' aExpr ']'
    ;

listElementIndexList
    :   listElementIndex
    |   listElementIndexList listElementIndex
    ;

listElementExpression
    :   Identifier listElementIndexList
    ;

assignInit
    :   aExpr
    |   listInitExpression
    |   funcInitExpression
    ;

assignInitList
    :   assignInit
    |   assignInitList ',' assignInit
    ;

assignStatement
    :   declaratorList '=' assignInitList
    ;

selectionStatement
    :   'if' bExp '{' funcStatementList? '}' ('elsif' bExp '{' funcStatementList? '}')* ('else' '{' funcStatementList? '}')?
    |   bExp '?' funcStatement ':' funcStatement
    ;

iterationStatement
    :   'for' bExp? '{' funcStatementList '}'
    |   'for' assignStatement? ';' bExp? ';' assignStatement? '{' funcStatementList '}' // todo
    ;

definition
    :   assignStatement
    |   funcDefinition
    ;

definitionList
    :   definition
    |   definitionList definition
    ;

file // todo: package import
    :   declaratorList
    ;

typeSpecifier
    :   simpleTypeSpecifier
    |   listTypeSpecifier
    |   funcTypeSpecifier
    ;

typeSpecifierList
    :   typeSpecifier
    |   typeSpecifierList ',' typeSpecifier
    ;

paraDeclaratorList
    :   Identifier
    |   paraDeclaratorList ',' Identifier
    ;

paraDeclaratorWithIdentity // todo: rename
    :   declaratorList typeSpecifier
    ;

paraDeclaratorWithIdentityList
    :   paraDeclaratorWithIdentity
    |   paraDeclaratorWithIdentityList ',' paraDeclaratorWithIdentity
    ;

funcTypeSpecifier
    :   'func' '(' paraDeclaratorWithIdentityList? ')' ('(' typeSpecifierList? ')')?
    ;

funcTypeSpecifierWithName
    :   'func' Identifier '(' paraDeclaratorWithIdentityList? ')' ('(' typeSpecifierList? ')')?
    ;

funcReturnPara
    :   aExpr
    |   bExp
    |   Identifier
    |   'nil'
    ;

funcReturnParaList
    :   funcReturnPara
    |   funcReturnParaList ',' funcReturnPara
    ;

funcStatement
    :   assignStatement
    |   selectionStatement
    |   iterationStatement
    |   'return' funcReturnParaList?
    ;

funcStatementList
    :   funcStatement
    |   funcStatementList funcStatement
    ;

funcBody
    :   funcStatementList?
    ;

funcInitExpression // when assign to a varient: function = func() {}
    :   funcTypeSpecifier '{' funcBody? '}'
    ;

funcDefinition // when define a method for a class or a static function: func function() {}
    :   funcTypeSpecifierWithName '{' funcBody? '}'
    ;

WS : [ \t\r\n]+ -> skip ; // skip spaces, tabs, newlines
