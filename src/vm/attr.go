/*
 * Copyright (C) 2023 The GDLang Team.
 *
 * This file is part of GDLang.
 *
 * GDLang is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * GDLang is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with GDLang.  If not, see <http://www.gnu.org/licenses/>.
 */

package vm

import "gdlang/lib/runtime"

type GDVMIdentSymbol struct {
	Ident  runtime.GDIdent
	Symbol *runtime.GDSymbol
}

func (p *GDVMProc) evalAGet(stack *runtime.GDSymbolStack) (runtime.GDObject, error) {
	isNilSafe, err := p.ReadBool()
	if err != nil {
		return nil, err
	}

	// Reads the expression and the attribute
	expr, attrErr := p.ReadAttributable(stack)

	// Continue reading the ident, after validating the expression
	ident, err := p.ReadIdent()
	if err != nil {
		return nil, err
	}

	if attrErr != nil && !isNilSafe {
		return nil, err
	} else if attrErr != nil && isNilSafe {
		stack.PushBuffer(runtime.GDZNil)
		return nil, nil
	}

	symbol, err := expr.GetAttr(ident)
	if err != nil && !isNilSafe {
		return nil, err
	} else if err != nil && isNilSafe {
		stack.PushBuffer(runtime.GDZNil)
		return nil, nil
	}

	attrObj := runtime.NewGDAttrIdObject(ident, symbol.Object, expr)
	stack.PushBuffer(attrObj)

	return nil, nil
}
