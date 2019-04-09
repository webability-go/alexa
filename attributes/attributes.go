package attributes

import (
  "fmt"
)

type Attributes map[string]interface{}

func (att *Attributes)Set(id string, val interface{}) {
  (*att)[id] = val
}

func (att *Attributes)Get(id string) interface{} {
  return (*att)[id]
}

func (att *Attributes)GetString(id string) string {
  return fmt.Sprint((*att)[id])
}

func (att *Attributes)GetInt(id string) int {
  if val, ok := (*att)[id]; ok {
    switch val.(type) {
      case int: return val.(int)
      case float64: return int(val.(float64))
      case bool: if val.(bool) {
        return 1
      } else {
        return 0
      }
    }
  }
  return 0
}

func (att *Attributes)GetBool(id string) bool {
  if val, ok := (*att)[id]; ok {
    switch val.(type) {
      case bool: return val.(bool)
      case int: return val.(int)!=0
      case float64: return val.(float64)!=0
    }
  }
  return false
}

func (att *Attributes)GetFloat(id string) float64 {
  if val, ok := (*att)[id]; ok {
    switch val.(type) {
      case float64: return val.(float64)
      case int: return float64(val.(int))
      case bool: if val.(bool) {
        return 1.0
      } else {
        return 0.0
      }
    }
  }
  return 0.0
}
