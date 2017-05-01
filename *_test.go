package main

import "testing"

func TestRepo(t *testing.T){
   var x Task

   if x = RepoFindTask(1); x == (Task{}){
      t.Error("Could not find ID")
   }
}
