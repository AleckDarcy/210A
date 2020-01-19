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
    ;

listTypeSpecifier
    :   listElementTypeSpecifier
    |   '[' ']' listTypeSpecifier
    ;

aExp
    :   IntegerLiteral
    |   FloatLiteral
    |   Identifier
    |   listElementExpression
    |   aExp ('*'|'/') aExp
    |   aExp ('+'|'-') aExp
    |   '(' aExp ')'
    ;

aExpList
    :   aExp
    |   aExpList ',' aExp
    ;

bExp
    :   aExp ('=='|'<'|'>'|'>='|'<='|'!=') aExp
    |   bExp ('=='|'&&'|'||'|'!=') bExp
    |   '!' bExp
    |   '(' bExp ')'
    ;

integerExpression
    :   IntegerLiteral
    |   Identifier
    ;

listElement
    :   '[' IntegerLiteral ']'
    ;

listElements
    :   listElement
    |   listElements listElement
    ;

listElementExpression
    :   Identifier listElements
    ;

tupleSizes
    :   integerExpression
    |   tupleSizes ',' integerExpression
    ;

assignStatement
    :   declaratorList simpleTypeSpecifier
    |   declaratorList '=' aExpList
    |   declarator '=' 'list' '(' listTypeSpecifier ',' '(' tupleSizes ')' ')'
    |   declarator '=' funcExpression
    ;

selectionStatement
    :   'if' bExp '{' funcStatementList? '}' ('elsif' bExp '{' funcStatementList? '}')* ('else' '{' funcStatementList? '}')?
    |   bExp '?' funcStatement ':' funcStatement
    ;

iterationStatement
    :   'for' bExp? '{' funcStatementList '}'
    |   'for' assignStatement ';' bExp? ';' assignStatement '{' funcStatementList '}' // todo
    ;

entry
    :   assignStatement
    |   funcExpressionWithName
    ;

entryList
    :   entry
    |   entryList entry
    ;

typeSpecifier
    :   simpleTypeSpecifier
    |   listTypeSpecifier
    |   funcSpecifier
    ;

typeSpecifierList
    :   typeSpecifier
    |   typeSpecifierList ',' typeSpecifier
    ;

typeSpecifierWithIdentity // todo: rename
    :   declaratorList typeSpecifier
    ;

typeSpecifierWithIdentityList
    :   typeSpecifierWithIdentity
    |   typeSpecifierWithIdentityList ',' typeSpecifierWithIdentity
    ;

funcSpecifier
    :   'func' '(' typeSpecifierWithIdentityList? ')' ('(' typeSpecifierList? ')')?
    ;

funcSpecifierWithName
    :   'func' Identifier '(' typeSpecifierWithIdentityList? ')' ('(' typeSpecifierList? ')')?
    ;

funcReturnPara
    :   aExp
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

funcExpression // when assign to a varient: function = func() {}
    :   funcSpecifier '{' funcBody? '}'
    ;

funcExpressionWithName // when define a method for a class or a static function: func function() {}
    :   funcSpecifierWithName '{' funcBody? '}'
    ;

WS : [ \t\r\n]+ -> skip ; // skip spaces, tabs, newlines
