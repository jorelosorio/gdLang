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

package ir

import (
	"bytes"
	"fmt"
	"gdlang/src/cpu"
	"gdlang/src/gd/ast"
)

type GDIRTIf struct {
	expr     GDIRNode
	thenExpr GDIRNode
	elseExpr GDIRNode
	GDIRBaseNode
}

func (t *GDIRTIf) BuildAssembly(padding string) string {
	return padding + fmt.Sprintf("%s %s %s %s", cpu.GetCPUInstName(cpu.Tif), t.expr.BuildAssembly(""), t.thenExpr.BuildAssembly(""), t.elseExpr.BuildAssembly(""))
}

func (t *GDIRTIf) BuildBytecode(bytecode *bytes.Buffer, ctx *GDIRContext) error {
	ctx.AddMapping(bytecode, t.GetPosition())

	err := Write(bytecode, cpu.Tif)
	if err != nil {
		return err
	}

	err = t.elseExpr.BuildBytecode(bytecode, ctx)
	if err != nil {
		return err
	}

	err = t.thenExpr.BuildBytecode(bytecode, ctx)
	if err != nil {
		return err
	}

	err = t.expr.BuildBytecode(bytecode, ctx)
	if err != nil {
		return err
	}

	return nil
}

func NewGDIRTIf(expr GDIRNode, thenExpr GDIRNode, elseExpr GDIRNode, node ast.Node) (*GDIRTIf, *GDIRObject) {
	return &GDIRTIf{expr, thenExpr, elseExpr, GDIRBaseNode{node}}, NewGDIRRegObject(cpu.RPop, node)
}
