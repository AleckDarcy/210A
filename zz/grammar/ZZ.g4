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

identifier
    :   Identifier
    ;

declarator
    :   identifier # declarator_identifier
    |   listElementExpression # declarator_listElementExpression
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
    :   'list' '(' listTypeSpecifier ',' aExpr ')'
    ;

aExpr
    :   IntegerLiteral # aExp_IntergerLiteral
    |   FloatLiteral # aExp_FloatLiteral
    |   identifier # aExp_identifier
    |   listElementExpression # aExp_listElementExpression
    |   aExpr ('*'|'/') aExpr # aExp_multiplicativeExpression
    |   aExpr ('+'|'-') aExpr # aExp_additiveExpression
    |   '(' aExpr ')' # aExp_bracketExpression
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

integerExpression
    :   IntegerLiteral
    |   identifier
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
    :   aExpr # assignInit_aExpr
    |   listInitExpression # assignInit_listInitExpression
    |   funcInitExpression # assignInit_funcInitExpression
    ;

assignInitList
    :   assignInit
    |   assignInitList ',' assignInit
    ;

assignStatement
    :   declaratorList '=' assignInitList
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

iterationStatement
    :   'for' bExpr? '{' funcStatementList '}'
    |   'for' assignStatement? ';' bExpr? ';' assignStatement? '{' funcStatementList '}' // todo
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
    :   definitionList
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
    :   identifier
    |   paraDeclaratorList ',' identifier
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
    |   selectionStatement
    |   iterationStatement
    |   funcReturnStatement
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
