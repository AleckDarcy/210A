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

bExpr
    :   aExpr ('=='|'<'|'>'|'>='|'<='|'!=') aExpr
    |   bExpr ('=='|'&&'|'||'|'!=') bExpr
    |   '!' bExpr
    |   '(' bExpr ')'
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
    |   funcExecuteExpression
    ;

assignInitList
    :   assignInit
    |   assignInitList ',' assignInit
    ;

assignStatement
    :   declaratorList '=' assignInitList
    ;

selectionStatement
    :   'if' bExpr '{' funcStatementList? '}' ('elsif' bExpr '{' funcStatementList? '}')* ('else' '{' funcStatementList? '}')?
    |   bExpr '?' funcStatement ':' funcStatement
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
    |   bExpr
    |   Identifier
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
    |   funcExecuteStatement
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

WS : [ \t\r\n]+ -> skip ; // skip spaces, tabs, newlines
